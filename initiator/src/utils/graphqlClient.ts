import { GraphQLClient } from "graphql-request";
import { HASURA_GRAPHQL_ADMIN_SECRET, HASURA_GRAPHQL_ENDPOINT } from "../../config";



export const graphqlClient = new GraphQLClient(
	HASURA_GRAPHQL_ENDPOINT,
	{
		headers:
		{
			"Content-Type": "application/json",
			"x-hasura-admin-secret": HASURA_GRAPHQL_ADMIN_SECRET,
		}
	} );
