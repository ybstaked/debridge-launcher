import { gql } from "graphql-request";



export const UPDATE_CHAIN_LATEST_BLOCK_MUTATION = gql`
	mutation UpdateChainLatestBlock( $chainId: bigint!, $latestBlock: Int! )
	{
		update_supported_chains_by_pk( pk_columns: {chainId: $chainId}, _set: {latestBlock: $latestBlock} )
		{
			latestBlock
		}
	}
`;
