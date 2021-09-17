
module.exports = {
	"schema":
	[
		{
			"http://localhost:8090/v1/graphql": {
				"headers": {
					"x-hasura-admin-secret": ""
				}
			}
		}
	],
	"documents": [
		"./src/**/*.ts*",
		"./src/**/*.js*"
	],
	"overwrite": true,
	"generates": {
		"./src/generated/graphql.ts": {
			"plugins": [
				"typescript",
				"typescript-operations",
//				"typescript-react-apollo"
			],
			"config": {
				"skipTypename": false,
				"withHooks": false,
				"withHOC": false,
				"withComponent": false
			}
		},
		"./graphql.schema.json": {
			"plugins": [
				"introspection"
			]
		}
	}
};
