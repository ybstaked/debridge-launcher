import { gql } from "graphql-request";



export const GET_CHAIN_SUBMISSION_DETAILS_QUERY = gql`
	query getChainSubmissionDetails( $chainIdTo: Int!, $submissionId: Int! )
	{
		aggregator_chains( where: { chainIdTo: { _eq: $chainIdTo } } )
		{
			supportedChain
			{
				chainlinkConfigs
				{
					chainId
				}
			}
		}

		submissions_by_pk( id: $submissionId )
		{
			id
		}
	}
`;
