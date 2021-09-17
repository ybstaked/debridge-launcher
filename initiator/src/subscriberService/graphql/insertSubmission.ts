import { gql } from "graphql-request";



export const INSERT_SUBMISSION_MUTATION = gql`
	mutation InsertSubmission( $object: submissions_insert_input! )
	{
		insert_submissions_one( object: $object )
		{
			id
		}
	}
`;
