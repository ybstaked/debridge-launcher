import { gql } from "graphql-request";



export const GET_CHAIN_DETAILS_QUERY = gql`
	query GetChainDetails( $chainId: Int! )
	{
		supported_chains_by_pk( chainId: $chainId )
		{
			chainId
			latestBlock
			chainlinkConfigs
			{
				eiChainlinkUrl
				eiCiAccesskey
				eiCiSecret
				eiIcAccesskey
				eiIcSecret
				provider
				debridgeAddr
				network
			}
		}
	}
`;
