import { gql } from "graphql-request";



export const GET_SUBMISSIONS_BY_CONFIRMATION_ID = gql`
	query GetCreatedSubmissionsByConfiramationChainId( $confirmationChainId: bigint! )
	{
		get_created_submissions_by_confirmation_chain_id(args: {in_confirmation_chain_id: $confirmationChainId})
		{
			submissionId
		}
	}
`;
