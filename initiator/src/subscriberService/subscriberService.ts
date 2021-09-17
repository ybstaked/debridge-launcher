import Web3 from "web3";
import { abi as whiteDebridgeAbi } from "../assets/WhiteFullDebridge.json";
// import { Chainlink } from './chainlink';
import { EventData } from "web3-eth-contract";
import { REQUIRE_MIN_CONFIRMATIONS } from '../../config';
import { GetChainDetailsQuery, GetChainDetailsQueryVariables, GetChainsConfigQuery, GetChainSubmissionDetailsQuery, GetChainSubmissionDetailsQueryVariables, InsertSubmissionMutationVariables } from "../generated/graphql";
import { graphqlClient } from '../utils/graphqlClient';
import { GET_CHAINS_CONFIG_QUERY } from './graphql/getChainsConfig';
import { GET_CHAIN_DETAILS_QUERY } from "./graphql/getChainDetails";
import { GET_CHAIN_SUBMISSION_DETAILS_QUERY } from "./graphql/getChainSubmissionDetails";
import { INSERT_SUBMISSION_MUTATION } from "./graphql/insertSubmission";


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

		// setInterval( async () =>
		// 	{
		// 		try
		// 		{
		// 			await checkConfirmations( chainConfig );
		// 		}
		// 		catch( exception )
		// 		{
		// 			console.error( exception );
		// 		}
		// 	}, interval );
	}

	console.info( "setInterval 864000 for setAllChainlinkCookies" );

	// setInterval(async () => {
	// 	try {
	// 		await this.setAllChainlinkCookies();
	// 	} catch (e) {
	// 		this.log.error(e);
	// 	}
	// }, 864000);
}



async function checkNewEvents( chainId: number )
{
	// this.logger.verbose(`checkNewEvents ${chainId}`);

	const payload: GetChainDetailsQueryVariables = { chainId };

	const chainsDetailsData: GetChainDetailsQuery = await graphqlClient.request( GET_CHAIN_DETAILS_QUERY, payload );

	const { supported_chains_by_pk: chainsDetails } = chainsDetailsData;

	// TODO: ****** Спросить у Ярослава, зачем две таблицы и есть ли там по факту связь один к многим *******
	if( !chainsDetails || !chainsDetails.chainlinkConfigs.length )
	{
		console.error( `Chain config not found by chainIn ${ chainId }` );
		return;
	}

	const chainConfig = chainsDetails.chainlinkConfigs[ 0 ];

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

	const isOk1 = await processNewTransfers( sentEvents, chainsDetails.chainId );
	const isOk2 = await processNewTransfers( burntEvents, chainsDetails.chainId );

	/* update lattest viewed block */
	//supportedChain.latestBlock = toBlock;
	if( isOk1 && isOk2 )
	{
		// TODO: ************** Дописать!!!!!!! **************************************
		// if( supportedChain[ LATEST_BLOCK_KEY ] != toBlock )
		// {
		// 	console.log( `updateSupportedChainBlock chainId: ${chainId}; key: ${ LATEST_BLOCK_KEY }; value: ${ toBlock }` );

		// 	await this.supportedChainRepository.update(chainId, {
		// 		latestBlock: toBlock,
		// 	});
		// }
	}
	else
	{
		console.error( `checkNewEvents. Last block not updated. Found error in processNewTransfers ${ chainId }` );
	}
}


async function processNewTransfers( events: EventData[], chainIdFrom: number )
{
	let result = true;

	if( !events )
		return result;

	// TODO: Переписать на вариант, когда мы один раз пробежимся по циклу,
	//  а потом одним запросом и одной мутацией всё обработаем группой
	for( const event of events )
	{
		console.log(`processNewTransfers chainIdFrom ${chainIdFrom}; submissionId: ${event.returnValues.submissionId}`);
		// this.logger.debug(event);

		const chainIdTo = event.returnValues.chainIdTo;
		const submissionId = event.returnValues.submissionId;

		let payload: GetChainSubmissionDetailsQueryVariables = { chainIdTo, submissionId };

		// !!!!!!!!!!!!!!!!!!!!!!!!!!!
		// Этот запрос делает что-то странное, он просто проверяет, что в таблице есть строки и всё
		// Спросить у Ярослава, если нас интересует только то, чтобы вставка была корректной, то я могу это сделать на уровне БД
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!
		const submissionDetails: GetChainSubmissionDetailsQuery = await graphqlClient.request( GET_CHAIN_SUBMISSION_DETAILS_QUERY, payload );

		const { submissions_by_pk: submission } = submissionDetails;

		// !!!!!!!!!!!!!!!!!!!!!!!!!!!
		// Уточнить согласно комментария выше и раскомментировать
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!
		// if (!chainConfig) {
		// 	this.logger.error(`Not found chainConfig: ${aggregatorInfo.aggregatorChain}`);
		// 	result = false;
		// 	continue;
		// }

		if( submission )
		{
			console.log( `Submission already found in db submissionId: ${ submissionId }` );
			continue;
		}


		// TODO: Переписать на групповую запись
		// Записываем задачу в БД, потом по отдельному планировщику все эти задачи будут обрабатываться пакетом

		// !!!!!!!!!!!!!!!!!!!!!!!!!!!
		// В таблице почему-то нет поля chainTo!!!!!!!!!!!!!
		// !!!!!!!!!!!!!!!!!!!!!!!!!!!
		const mutationPayload: InsertSubmissionMutationVariables =
		{
			object:
			{
				submissionId,
				txHash: "NULL",
				runId: null,
				chainFrom: chainIdFrom,
				// В таблице почему-то нет поля chainTo!!!!!!!!!!!!!
				// chainTo: event.returnValues.chainIdTo,
				debridgeId: event.returnValues.debridgeId,
				receiverAddress: event.returnValues.receiver,
				amount: event.returnValues.amount,
				status: SubmisionStatus.NEW
			}
		};

		await graphqlClient.request( INSERT_SUBMISSION_MUTATION, mutationPayload );
	}

	return result;
}

// class Subscriber {
// 	db: Db;
// 	chainlink: Chainlink;
// 	log: log4js.Logger;

// 	constructor() {
// 		this.db = new Db();
// 		this.chainlink = new Chainlink();
// 		this.log = log4js.getLogger('subscriber');
// 	}

// 	async init() {
// 		this.log.info('init');
// 		await this.db.connectDb();
// 		await this.db.createTables();
// 		await this.setAllChainlinkCookies();
// 		await this.subscribe();
// 	}

// 	/* call the chainlink node and run a job */
// 	async subscribe() {
// 		const supportedChains = await this.db.getSupportedChains();
// 		for (const supportedChain of supportedChains) {
// 			//const web3 = new Web3(supportedChain.provider);
// 			//const registerInstance = new web3.eth.Contract(
// 			//    whiteDebridgeAbi,
// 			//    supportedChain.debridgeaddr
// 			//);

// 			this.log.info(`setInterval ${supportedChain.interval} for checkNewEvents ${supportedChain.network}`);
// 			setInterval(async () => {
// 				try {
// 					await this.checkNewEvents(supportedChain.chainId);
// 				} catch (e) {
// 					this.log.error(e);
// 				}
// 			}, supportedChain.interval);
// 		}
// 		const chainConfigs = await this.db.getChainConfigs();
// 		for (const chainConfig of chainConfigs) {
// 			this.log.info(`setInterval 30000 for checkConfirmations ${chainConfig.network}`);
// 			setInterval(async () => {
// 				try {
// 					await this.checkConfirmations(chainConfig);
// 				} catch (e) {
// 					this.log.error(e);
// 				}
// 			}, 30000);
// 		}

// 		this.log.info(`setInterval 864000 for setAllChainlinkCookies`);
// 		setInterval(async () => {
// 			try {
// 				await this.setAllChainlinkCookies();
// 			} catch (e) {
// 				this.log.error(e);
// 			}
// 		}, 864000);
// 	}

// 	/* collect new events */
// 	async checkNewEvents(chainId: number) {
// 		this.log.debug(`checkNewEvents ${chainId}`);
// 		const supportedChain = await this.db.getSupportedChain(chainId);

// 		const web3 = new Web3(supportedChain.provider);
// 		const registerInstance = new web3.eth.Contract(whiteDebridgeAbi as any, supportedChain.debridgeAddr);
// 		/* get blocks range */
// 		//console.log(await web3.eth.getBlockNumber());
// 		const toBlock = (await web3.eth.getBlockNumber()) - MIN_CONFIRMATIONS;
// 		const fromBlock = supportedChain.latestBlock > 0 ? supportedChain.latestBlock : toBlock - 1;

// 		if (fromBlock >= toBlock) return;
// 		this.log.info(`checkNewEvents ${supportedChain.network} ${fromBlock}-${toBlock}`);

// 		/* get events */
// 		const sentEvents = await registerInstance.getPastEvents(
// 			'Sent',
// 			{ fromBlock, toBlock }, //,
// 			//async (error, events) => {
// 			//    if (error) {
// 			//        this.log.error(error);
// 			//    }
// 			//    await this.processNewTransfers(events, supportedChain.chainId);
// 			//}
// 		);
// 		const burntEvents = await registerInstance.getPastEvents(
// 			'Burnt',
// 			{ fromBlock, toBlock }, //,
// 			//async (error, events) => {
// 			//    if (error) {
// 			//        this.log.error(error);
// 			//    }
// 			//await this.processNewTransfers(events, supportedChain.chainId);
// 			//}
// 		);

// 		const isOk1 = await this.processNewTransfers(sentEvents, supportedChain.chainId);
// 		const isOk2 = await this.processNewTransfers(burntEvents, supportedChain.chainId);

// 		/* update lattest viewed block */
// 		//supportedChain.latestBlock = toBlock;
// 		if (isOk1 && isOk2) {
// 			await this.db.updateSupportedChainKey(supportedChain.chainId, 'latestBlock', toBlock);
// 		} else {
// 			this.log.error(`checkNewEvents. Last block not updated. Found error in processNewTransfers ${chainId}`);
// 		}
// 	}

// 	/* proccess new events */
// 	async processNewTransfers(events: EventData[], chainIdFrom: number) {
// 		if (!events) return true;
// 		let isOk = true;
// 		for (const e of events) {
// 			this.log.info(`processNewTransfers chainIdFrom ${chainIdFrom}; submissionId: ${e.returnValues.submissionId}`);
// 			this.log.debug(e);
// 			/* remove chainIdTo  selector */
// 			const chainIdTo = e.returnValues.chainIdTo;
// 			const aggregatorInfo = await this.db.getAggregatorConfig(chainIdTo);
// 			if (!aggregatorInfo) continue;
// 			const chainConfig = await this.db.getChainConfig(aggregatorInfo.aggregatorChain);
// 			if (!chainConfig) {
// 				this.log.error(`Not found chainConfig: ${aggregatorInfo.aggregatorChain}`);
// 				isOk = false;
// 				continue;
// 			}

// 			/* call chainlink node */
// 			const submissionId = e.returnValues.submissionId;
// 			const submission = await this.db.getSubmission(submissionId);
// 			if (submission) {
// 				this.log.debug(`Submission already found in db submissionId: ${submissionId}`);
// 				continue;
// 			}
// 			await this.callChainlinkNode(chainConfig.mintJobId, chainConfig, submissionId, e.returnValues, chainIdFrom);
// 		}
// 		return isOk;
// 	}

// 	/* set chainlink cookies */
// 	async checkConfirmations(chainConfig) {
// 		//this.log.debug(`checkConfirmations ${chainConfig.network}`);

// 		//get chainTo for current aggregator
// 		const supportedChains = await this.db.getChainToForAggregator(chainConfig.chainId);
// 		this.log.debug(`checkConfirmations ${chainConfig.network} check submissions to network ${supportedChains.map(a => a.chainTo)}`);
// 		const createdSubmissions = await this.db.getSubmissionsByStatusAndChainTo(
// 			SubmisionStatus.CREATED,
// 			supportedChains.map(a => a.chainTo),
// 		);
// 		for (const submission of createdSubmissions) {
// 			const runInfo = await this.chainlink.getChainlinkRun(chainConfig.eichainlinkurl, submission.runId, chainConfig.cookie);
// 			if (runInfo) {
// 				if (runInfo.status == 'completed') await this.db.updateSubmissionStatus(submission.submissionId, SubmisionStatus.CONFIRMED);
// 				if (runInfo.status == 'errored') await this.db.updateSubmissionStatus(submission.submissionId, SubmisionStatus.REVERTED);
// 			}
// 		}
// 	}

// 	/* call the chainlink node and run a job */
// 	async callChainlinkNode(jobId: string, chainConfig, submissionId: string, e, chainIdFrom: number) {
// 		this.log.info(`callChainlinkNode jobId ${jobId}; submissionId: ${submissionId}`);
// 		const runId = await this.chainlink.postChainlinkRun(
// 			jobId,
// 			submissionId,
// 			chainConfig.eichainlinkurl,
// 			chainConfig.eiicaccesskey,
// 			chainConfig.eiicsecret,
// 		);
// 		this.log.info(`Received runId ${runId}; submissionId: ${submissionId}`);
// 		await this.db.createSubmission({
// 			submissionId,
// 			txHash: 'NULL',
// 			runId,
// 			chainFrom: chainIdFrom,
// 			chainTo: e.chainIdTo,
// 			debridgeId: e.debridgeId,
// 			receiverAddr: e.receiver,
// 			amount: e.amount,
// 			status: SubmisionStatus.CREATED,
// 		});
// 	}

// 	/* set chainlink cookies */
// 	async setAllChainlinkCookies() {
// 		this.log.debug(`Start setAllChainlinkCookies`);
// 		const chainConfigs = await this.db.getChainConfigs();
// 		for (const chainConfig of chainConfigs) {
// 			this.log.debug(`setAllChainlinkCookies ${chainConfig.network}`);
// 			const cookies = await this.chainlink.getChainlinkCookies(chainConfig.eiChainlinkUrl, chainConfig.network);
// 			await this.db.updateChainConfigCookie(chainConfig.chainId, cookies);
// 		}
// 	}
// }

// export { Subscriber };
