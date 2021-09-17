import fetch from "node-fetch";



// ***************************************************************************************************
// ***************************************************************************************************
export async function postChainlinkBatchRun(
	jobId: string,
	data: Array<string>,
	eiChainlinkUrl: string,
	eiIcAccessKey: string,
	eiIcSecret: string
)
{
	// Метод не реализован на сервере
	return "123e4567-e89b-12d3-a456-426655440000";


	// const postJobUrl = eiChainlinkUrl + "/v2/specs/" + jobId + "/runs";

	// console.log( "postChainlinkRun", data );

	// const body = { result: data };

	// const settings =
	// {
	// 	method: "POST",
	// 	headers:
	// 	{
	// 		"Content-Type": "application/json",
	// 		"X-Chainlink-EA-AccessKey": eiIcAccessKey,
	// 		"X-Chainlink-EA-Secret": eiIcSecret,
	// 	},
	// 	body: JSON.stringify( body )
	// };

	// const response = await fetch( postJobUrl, settings );

	// const responseData: ChainlinkPostResponseType = await response.json();

	// console.log( "postChainlinkRun response", responseData.data.data );

	// return responseData.data.data.id;
}



// ***************************************************************************************************
// ***************************************************************************************************
export async function postChainlinkRun(
	jobId: string,
	data: string,
	eiChainlinkUrl: string,
	eiIcAccessKey: string,
	eiIcSecret: string
)
{
	const postJobUrl = eiChainlinkUrl + "/v2/specs/" + jobId + "/runs";

	console.log( "postChainlinkRun", data );

	const body = { result: data };

	const settings =
	{
		method: "POST",
		headers:
		{
			"Content-Type": "application/json",
			"X-Chainlink-EA-AccessKey": eiIcAccessKey,
			"X-Chainlink-EA-Secret": eiIcSecret,
		},
		body: JSON.stringify( body )
	};

	const response = await fetch( postJobUrl, settings );

	const responseData: ChainlinkPostResponseType = await response.json();

	console.log( "postChainlinkRun response", responseData.data.data );

	return responseData.data.data.id;
}



// ***************************************************************************************************
// ***************************************************************************************************
export async function getChainlinkRun( eiChainlinkUrl: string, runId: string, cookie: string )
{
    const getRunUrl = '/v2/runs/' + runId;

	const settings =
	{
		method: "GET",
		headers: { "Content-Type": "application/json" },
		Cookie: JSON.parse(cookie)
	};

	const response = await fetch( getRunUrl, settings );

	const responseData: ChainlinkPostResponseType = await response.json();

	return responseData.data.data.attributes;
}



// ***************************************************************************************************
// ***************************************************************************************************
interface ChainlinkPostResponseType
{
	data:
	{
		data:
		{
			id: string,
			attributes: any
		}
	}
}


