import { gql } from "graphql-request";



export const UPDATE_SUBMISSIONS_MUTATION = gql`
	mutation UpdateSubmissions( $runId: String!, $status: Int!, $submissionIds: [String!]! )
	{
		update_submissions( where: {submissionId: {_in: $submissionIds}}, _set: {runId: $runId, status: $status} )
		{
			affected_rows
		}
	}
`;
