import { gql } from "graphql-request";


export const GET_CHAINS_CONFIG_QUERY = gql`
	query GetChainsConfig
	{
		supported_chains
		{
			chainId
		}

		chainlink_configs
		{
			chainId
			network
		}
	}
`;
