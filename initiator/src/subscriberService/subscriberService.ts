import Web3 from "web3";
import { abi as whiteDebridgeAbi } from "../assets/WhiteFullDebridge.json";
// import { Chainlink } from './chainlink';
import { EventData } from "web3-eth-contract";
import { PUT_JOBS_TO_CHAINLINK_INTERVAL, REQUIRE_MIN_CONFIRMATIONS } from '../../config';
import { GetChainDetailsQuery, GetChainDetailsQueryVariables, GetChainsConfigQuery, GetChainSubmissionDetailsQuery, GetChainSubmissionDetailsQueryVariables, GetCreatedSubmissionsByConfiramationChainIdQuery, GetCreatedSubmissionsByConfiramationChainIdQueryVariables, GetNewSubmissionsQuery, InsertSubmissionMutationVariables, UpdateChainLatestBlockMutationVariables, UpdateSubmissionsMutation, UpdateSubmissionsMutationVariables } from "../generated/graphql";
import { graphqlClient } from '../utils/graphqlClient';
import { GET_CHAINS_CONFIG_QUERY } from './graphql/getChainsConfig';
import { GET_CHAIN_DETAILS_QUERY } from "./graphql/getChainDetails";
import { GET_CHAIN_SUBMISSION_DETAILS_QUERY } from "./graphql/getChainSubmissionDetails";
import { INSERT_SUBMISSION_MUTATION } from "./graphql/insertSubmission";
import { GET_NEW_SUBMISSIONS_QUERY } from "./graphql/getNewSubmission";
import groupBy from "lodash/groupBy";
import { getChainlinkRun, postChainlinkBatchRun, postChainlinkRun } from "./chainlinkConnector/chainlinkConnector";
import { UPDATE_SUBMISSIONS_MUTATION } from "./graphql/updateSubmissions";
import { UPDATE_CHAIN_LATEST_BLOCK_MUTATION } from "./graphql/updateChainLatestBlock";
import { GET_SUBMISSIONS_BY_CONFIRMATION_ID } from "./graphql/getSubmissionsByConfirmationChainId";


const MIN_CONFIRMATIONS: number = REQUIRE_MIN_CONFIRMATIONS; // parseInt( process.env.MIN_CONFIRMATIONS || 3 );

const LATEST_BLOCK_KEY = "latestBlock";

const SubmisionStatus =
{
	NEW: 0, // Job accumulated, but not sent to chainlink node
	CREATED: 1, // Job sent to chainlink node
	BROADCASTED: 2,
	CONFIRMED: 3,
	REVERTED: 4
};



// ***************************************************************************************************
// ***************************************************************************************************
export async function subscribeToBlockchainEvents()
{
	const chainsConfig: GetChainsConfigQuery = await graphqlClient.request( GET_CHAINS_CONFIG_QUERY );

	const { supported_chains: supportedChains, chainlink_configs: chainConfigs } = chainsConfig;


	for( const supportedChain of supportedChains )
	{
		const interval = 3000; // supportedChain.interval

		console.info( `setInterval ${ interval } for checkNewEvents ${ supportedChain.network }` );

		setInterval( async () =>
			{
				try
				{
					await checkNewEvents( supportedChain.chainId );
				}
				catch( exception )
				{
					console.error( exception );
				}
			}, interval );
	}


	for( const chainConfig of chainConfigs )
	{
		const interval = 30000;

		console.info( `setInterval ${ interval } for checkConfirmations ${ chainConfig.network }` );

		setInterval( async () =>
			{
				try
				{
					await checkConfirmations( chainConfig );
				}
				catch( exception )
				{
					console.error( exception );
				}
			}, interval );
	}

	console.info( "setInterval 864000 for setAllChainlinkCookies" );

	// setInterval(async () => {
	// 	try {
	// 		await this.setAllChainlinkCookies();
	// 	} catch (e) {
	// 		this.log.error(e);
	// 	}
	// }, 864000);


	// Собираем из базы все события со статусом NEW и отправляем их на chainlink
	setInterval( async () =>
		{
			try
			{
				await batchChainlinkJobs();
			} catch (e) {
				this.logger.error(e);
			}
		}, PUT_JOBS_TO_CHAINLINK_INTERVAL );
}



// ***************************************************************************************************
// ***************************************************************************************************
async function checkNewEvents( chainId: number )
{
	console.log( `checkNewEvents ${ chainId }` );

	const payload: GetChainDetailsQueryVariables = { chainId };

	const chainsDetailsData: GetChainDetailsQuery = await graphqlClient.request( GET_CHAIN_DETAILS_QUERY, payload );

	const { supported_chains_by_pk: chainsDetails } = chainsDetailsData;

	if( !chainsDetails || !chainsDetails.config )
	{
		console.error( `Chain config not found by chainIn ${ chainId }` );
		return;
	}

	const chainConfig = chainsDetails.config;

	const web3 = new Web3( chainConfig.provider );
	const registerInstance = new web3.eth.Contract( whiteDebridgeAbi as any, chainConfig.debridgeAddr );

	// Get blocks range
	const toBlock = ( await web3.eth.getBlockNumber() ) - MIN_CONFIRMATIONS;
	const fromBlock = chainsDetails.latestBlock > 0 ? chainsDetails.latestBlock : toBlock - 1;

	if( fromBlock >= toBlock )
		return;

	console.log( `checkNewEvents ${ chainConfig.network } ${ fromBlock }-${ toBlock }` );

	/* get events */
	const sentEvents = await registerInstance.getPastEvents( "Sent", { fromBlock, toBlock } );
	const burntEvents = await registerInstance.getPastEvents( "Burnt", { fromBlock, toBlock } );

	// TODO: Объединить события в один массив или переделать на промис
	await processNewTransfers( sentEvents, chainsDetails.chainId );
	await processNewTransfers( burntEvents, chainsDetails.chainId );

	/* update lattest viewed block */
	if( chainsDetails.latestBlock != toBlock )
	{
		console.log( `updateSupportedChainBlock chainId: ${chainId}; key: ${ LATEST_BLOCK_KEY }; value: ${ toBlock }` );

		let mutationPayload: UpdateChainLatestBlockMutationVariables =
		{
			chainId,
			latestBlock:toBlock
		};

		graphqlClient.request( UPDATE_CHAIN_LATEST_BLOCK_MUTATION, mutationPayload );
	}
}



// ***************************************************************************************************
// ***************************************************************************************************
async function processNewTransfers( events: EventData[], chainIdFrom: number )
{
	// TODO: Переписать на вариант, когда мы один раз пробежимся по циклу,
	//  а потом одним запросом и одной мутацией всё обработаем группой
	for( const event of events )
	{
		console.log(`processNewTransfers chainIdFrom ${chainIdFrom}; submissionId: ${event.returnValues.submissionId}`);
		// this.logger.debug(event);

		// Проверим, что у нас есть конфигурация для этого чейна и что такого сабмишна уже нет в базе
		const confirmationChainId = event.returnValues.chainIdTo;
		const submissionId = event.returnValues.submissionId;

		const payload: GetChainSubmissionDetailsQueryVariables = { confirmationChainId, submissionId };

		const submissionDetails: GetChainSubmissionDetailsQuery = await graphqlClient.request( GET_CHAIN_SUBMISSION_DETAILS_QUERY, payload );

		const { submissions_by_pk: submission } = submissionDetails;

		if( submission )
		{
			console.log( `Submission already found in db submissionId: ${ submissionId }` );
			continue;
		}


		// TODO: Переписать на групповую запись
		// Записываем задачу в БД, потом по отдельному планировщику все эти задачи будут обрабатываться пакетом
		const mutationPayload: InsertSubmissionMutationVariables =
		{
			object:
			{
				submissionId,
				txHash:				"NULL",
				runId:				null,
				chainFrom:			chainIdFrom,
				chainTo:			event.returnValues.chainIdTo,
				debridgeId:			event.returnValues.debridgeId,
				receiverAddress:	event.returnValues.receiver,
				amount:				event.returnValues.amount,
				status:				SubmisionStatus.NEW
			}
		};

		await graphqlClient.request( INSERT_SUBMISSION_MUTATION, mutationPayload );
	}
}



// ***************************************************************************************************
// ***************************************************************************************************
async function batchChainlinkJobs()
{
	// Выборку надо группировать по chainTo и отправлять одним пакетом именно в рамках одной chainTo
	const newSubmissions: GetNewSubmissionsQuery = await graphqlClient.request( GET_NEW_SUBMISSIONS_QUERY );

	if( !newSubmissions.submissions.length || !newSubmissions.chainlink_configs )
		return;

	// Сконвертируем массив конфигураций в Map для скорости
	const configsMap = new Map( newSubmissions.chainlink_configs.map( ( item ) => [ item.chainId, item ] ) );
	//transofmConfigsToMap( newSubmissions.chainlink_configs );


	// Сгруппируем сабмишны по полю chainTo
	const groupedSubmissions = groupBy( newSubmissions.submissions, "chainTo" );

	for( const [ chainTo, submissionGroup ] of Object.entries( groupedSubmissions ) )
	{
		const chainConfig = configsMap.get( chainTo );

		if( !chainConfig )
			continue;

		let jobId: string = "";
		let runId: string = "";

		if( submissionGroup.length > 1 )
		{
			// Если у нас для chainlik несколько задач, то используем submitManyJobId
			jobId = chainConfig.submitManyJobId;

			runId =
				await postChainlinkBatchRun( jobId,
					submissionGroup.map( ( item ) => item.submissionId ),
					chainConfig.eiChainlinkUrl,
					chainConfig.eiIcAccesskey,
					chainConfig.eiIcSecret
				);
		}
		else
		{
			// Если у нас для chainlik только одна задача, то используем submitJobId
			jobId = chainConfig.submitJobId;

			runId =
				await postChainlinkRun(
					jobId,
					submissionGroup[ 0 ].submissionId,
					chainConfig.eiChainlinkUrl,
					chainConfig.eiIcAccesskey,
					chainConfig.eiIcSecret
				);
		}

		// TODO: Теперь нужно присвоить статус CREATED и прописать runId
		let payload: UpdateSubmissionsMutationVariables =
		{
			runId,
			status: SubmisionStatus.CREATED,
			submissionIds: submissionGroup.map( ( item ) => item.submissionId )
		};

		graphqlClient.request( UPDATE_SUBMISSIONS_MUTATION, payload );
	}
}



// ***************************************************************************************************
// ***************************************************************************************************
async function checkConfirmations( chainConfig )
{
	console.log( `checkConfirmations ${ chainConfig.network }` );

	const payload: GetCreatedSubmissionsByConfiramationChainIdQueryVariables =
	{
		confirmationChainId: chainConfig.chainId
	};

	const queryResult: GetCreatedSubmissionsByConfiramationChainIdQuery =
		await graphqlClient.request( GET_SUBMISSIONS_BY_CONFIRMATION_ID, payload );

	const createdSubmissions = queryResult.get_created_submissions_by_confirmation_chain_id;

	for( const submission of createdSubmissions )
	{
		const runInfo = await getChainlinkRun( chainConfig.eiChainlinkUrl, submission.runId, chainConfig.cookie );

		if( !runInfo )
			continue;

		// if( runInfo.status == 'completed' )
		// {
		// 	await this.submissionsRepository.update(submission.submissionId, {
		// 		status: SubmisionStatusEnum.CONFIRMED,
		// 	});
		// }

		// if( runInfo.status == 'errored' )
		// {
		// 	await this.submissionsRepository.update(submission.submissionId, {
		// 		status: SubmisionStatusEnum.REVERTED,
		// 	});
		// }
	}




	// // get chainTo for current aggregator
	// const supportedChains = await this.aggregatorChainsRepository.find({
	// 	aggregatorChain: chainConfig.chainId,
	// });

	// console.log( `checkConfirmations ${ chainConfig.network } check submissions to network ${ supportedChains.map( a => a.chainIdTo ) }` );

	// const createdSubmissions = await this.submissionsRepository.find({
	// 	where: {
	// 	status: SubmisionStatusEnum.CREATED,
	// 	chainTo: In(supportedChains.map(a => a.chainIdTo)),
	// 	},
	// });

	// for( const submission of createdSubmissions )
	// {
	// 	const runInfo = await this.chainlinkService.getChainlinkRun(chainConfig.eiChainlinkUrl, submission.runId, chainConfig.cookie);

	// 	if( !runInfo )
	// 		continue;

	// 	if( runInfo.status == 'completed' )
	// 	{
	// 		await this.submissionsRepository.update(submission.submissionId, {
	// 			status: SubmisionStatusEnum.CONFIRMED,
	// 		});
	// 	}

	// 	if( runInfo.status == 'errored' )
	// 	{
	// 		await this.submissionsRepository.update(submission.submissionId, {
	// 			status: SubmisionStatusEnum.REVERTED,
	// 		});
	// 	}
	// }
}
