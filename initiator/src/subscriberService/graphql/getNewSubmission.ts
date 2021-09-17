import { gql } from "graphql-request";


export const GET_NEW_SUBMISSIONS_QUERY = gql`
	query GetNewSubmissions
	{
		submissions(where: {status: {_eq: 0}}, order_by: {chainTo: asc})
		{
			submissionId
			chainTo
		}

		chainlink_configs
		{
			chainId
			submitJobId
			submitManyJobId
			confirmNewAssetJobId
			eiChainlinkUrl
			eiIcAccesskey
			eiIcSecret
		}
	}
`;
