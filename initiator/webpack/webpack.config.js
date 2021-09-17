const path = require( "path" );
const nodeExternals = require( "webpack-node-externals" );
const ForkTsCheckerWebpackPlugin = require( 'fork-ts-checker-webpack-plugin' );


module.exports =
{
	name: "intiator",

	mode: "production",

	entry: "./src/index.ts",
	target: "node",
	externals: [ nodeExternals() ],

	watch: false,

	output:
	{
		filename: "index.js",
		publicPath: "/",
		path: path.resolve(__dirname, "../build"),
		clean: true
	},

	resolve:
	{
		extensions: [ ".ts", ".js" ],
	},

	module:
	{
		rules:
		[
			{
				test: /\.ts?$/,
				loader: "ts-loader",
				exclude: /node_modules/,
				options:
				{
					transpileOnly: true
				}
			}
		]
	},

	plugins: [
        new ForkTsCheckerWebpackPlugin() // run TSC on a separate thread
    ],

	optimization:
	{
		nodeEnv: "production",
		concatenateModules: true,
		usedExports: true,
		innerGraph: true,
		mangleExports: true,
		minimize: true,
	}
};
