
export const REQUIRE_MIN_CONFIRMATIONS = 3;
export const PUT_JOBS_TO_CHAINLINK_INTERVAL = 3000;


export const APPLICATION_PORT = process.env.PORT || 8080;

export const HASURA_GRAPHQL_JWT_SECRET =
{
	"type": "HS256",
	"key": "uYgRivH1bDbZ6xHvoGQY7mhZ1SOEHqIfsWN1dMvRm4ZLtXOAkAmuwyjVAeekH34kHlVpJ9"
};

export const HASURA_GRAPHQL_ENDPOINT = "http://localhost:8090/v1/graphql";
export const HASURA_GRAPHQL_ADMIN_SECRET = "";

export const LOG_TO_SYSLOG = false;
export const SYSLOG_PORT = 514;
export const SYSLOG_HOST = ""; // "10.10.3.177";
