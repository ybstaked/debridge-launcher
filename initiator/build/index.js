(()=>{"use strict";var e={n:i=>{var n=i&&i.__esModule?()=>i.default:()=>i;return e.d(n,{a:n}),n},d:(i,n)=>{for(var t in n)e.o(n,t)&&!e.o(i,t)&&Object.defineProperty(i,t,{enumerable:!0,get:n[t]})},o:(e,i)=>Object.prototype.hasOwnProperty.call(e,i)};((e,i,n)=>{require("core-js/stable"),require("regenerator-runtime/runtime");const t=require("express");var s=n.n(t);const o=require("stream");var c=n.n(o);const r=require("syslog");var a=n.n(r);process.env.PORT;var d,u,l,h,f,I,k,b,C,_,m,y,v;require("web3"),require("web3-eth-contract"),function(e){e.PkC0a689f14599492539c0587c06b="PK_c0a689f14599492539c0587c06b",e.UqC0a689f14599492539c0587c06b="UQ_c0a689f14599492539c0587c06b"}(d||(d={})),function(e){e.AggregatorChain="aggregatorChain",e.ChainIdTo="chainIdTo"}(u||(u={})),function(e){e.AggregatorChain="aggregatorChain",e.ChainIdTo="chainIdTo"}(l||(l={})),function(e){e.ChainlinkConfigsPkey="chainlink_configs_pkey"}(h||(h={})),function(e){e.ChainId="chainId",e.ConfirmNewAssetJobId="confirmNewAssetJobId",e.Cookie="cookie",e.EiChainlinkUrl="eiChainlinkUrl",e.EiCiAccesskey="eiCiAccesskey",e.EiCiSecret="eiCiSecret",e.EiIcAccesskey="eiIcAccesskey",e.EiIcSecret="eiIcSecret",e.Id="id",e.Network="network",e.SubmitJobId="submitJobId",e.SubmitManyJobId="submitManyJobId"}(f||(f={})),function(e){e.ChainId="chainId",e.ConfirmNewAssetJobId="confirmNewAssetJobId",e.Cookie="cookie",e.EiChainlinkUrl="eiChainlinkUrl",e.EiCiAccesskey="eiCiAccesskey",e.EiCiSecret="eiCiSecret",e.EiIcAccesskey="eiIcAccesskey",e.EiIcSecret="eiIcSecret",e.Id="id",e.Network="network",e.SubmitJobId="submitJobId",e.SubmitManyJobId="submitManyJobId"}(I||(I={})),function(e){e.Asc="asc",e.AscNullsFirst="asc_nulls_first",e.AscNullsLast="asc_nulls_last",e.Desc="desc",e.DescNullsFirst="desc_nulls_first",e.DescNullsLast="desc_nulls_last"}(k||(k={})),function(e){e.SubmissionsPkey="submissions_pkey"}(b||(b={})),function(e){e.Amount="amount",e.ChainFrom="chain_from",e.CreatedAt="created_at",e.DebridgeId="debridge_id",e.Id="id",e.ReceiverAddress="receiver_address",e.RunId="run_id",e.Status="status",e.SubmissionId="submission_id",e.TxHash="tx_hash"}(C||(C={})),function(e){e.Amount="amount",e.ChainFrom="chain_from",e.CreatedAt="created_at",e.DebridgeId="debridge_id",e.Id="id",e.ReceiverAddress="receiver_address",e.RunId="run_id",e.Status="status",e.SubmissionId="submission_id",e.TxHash="tx_hash"}(_||(_={})),function(e){e.PkE27d3238e9e6b29429b425ff993="PK_e27d3238e9e6b29429b425ff993",e.UqE27d3238e9e6b29429b425ff993="UQ_e27d3238e9e6b29429b425ff993"}(m||(m={})),function(e){e.ChainId="chainId",e.LatestBlock="latestBlock",e.Network="network"}(y||(y={})),function(e){e.ChainId="chainId",e.LatestBlock="latestBlock",e.Network="network"}(v||(v={}));const g=require("graphql-request"),p=new g.GraphQLClient("http://localhost:8090/v1/graphql",{headers:{"Content-Type":"application/json","x-hasura-admin-secret":""}}),w=g.gql`
	query GetChainsConfig
	{
		supported_chains
		{
			chainId
			network
			latestBlock
		}

		chainlink_configs
		{
			chainId
			network
		}
	}
`,A=g.gql`
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
			}
		}
	}
`;var S=function(e,i,n,t){return new(n||(n=Promise))((function(s,o){function c(e){try{a(t.next(e))}catch(e){o(e)}}function r(e){try{a(t.throw(e))}catch(e){o(e)}}function a(e){var i;e.done?s(e.value):(i=e.value,i instanceof n?i:new n((function(e){e(i)}))).then(c,r)}a((t=t.apply(e,i||[])).next())}))};function q(e){return S(this,void 0,void 0,(function*(){const i={chainId:e};yield p.request(A,i)}))}process.env.PORT;const E=s()(),N=new(c().Writable);a().createClient(514,""),N._write=(e,i,n)=>{console.log(e.toString()),n()},function(){S(this,void 0,void 0,(function*(){const e=yield p.request(w),{supported_chains:i,chainlink_configs:n}=e;for(const e of i){const i=3e3;console.info(`setInterval ${i} for checkNewEvents ${e.network}`),setInterval((()=>S(this,void 0,void 0,(function*(){try{yield q(e.chainId)}catch(e){console.error(e)}}))),i)}for(const e of n){const i=3e4;console.info(`setInterval ${i} for checkConfirmations ${e.network}`)}console.info("setInterval 864000 for setAllChainlinkCookies")}))}(),E.use((function(e,i,n,t){var s;if(e)return N.write(e.message),N.write(e.stack),n.status((null===(s=null==e?void 0:e.output)||void 0===s?void 0:s.statusCode)||500).json(e.output.payload)}))})(0,0,e)})();