import "core-js/stable";
import "regenerator-runtime/runtime";
import express from "express";
import stream from "stream";
import syslog from "syslog";
import { APPLICATION_PORT, LOG_TO_SYSLOG, SYSLOG_HOST, SYSLOG_PORT } from "../config";
import { subscribeToBlockchainEvents } from "./subscriberService/subscriberService";


const PORT = process.env.PORT || APPLICATION_PORT || 8080;

const app = express();

const logStream = new stream.Writable();

// Сделаем отправку вывода на syslog сервер
const syslogLogger = syslog.createClient( SYSLOG_PORT, SYSLOG_HOST );

// Сделаем свой поток для логгирования
logStream._write = ( chunk, encoding, done ) =>
{
	if( LOG_TO_SYSLOG )
		syslogLogger.info( chunk.toString() );

	console.log( chunk.toString() );
	done();
};

// middleware
// app.use( express.json() );

// logger
// app.use( morgan( "tiny", { stream: logStream } ) );

// app.use( cookieParser() );

// Main cycle
subscribeToBlockchainEvents();

// Роут до обработчика
// app.use( "/",  );

// error handler
app.use( errorHandler );

// app.listen( PORT, () => logStream.write( "*********************** Listening on port " + PORT ) );



function errorHandler( error, request, response, next )
{
	if( !error )
		return;

	logStream.write( error.message );
	logStream.write( error.stack );

	return response.status( error?.output?.statusCode || 500 ).json( error.output.payload );
}

















// process.env['NODE_CONFIG_DIR'] = __dirname + '/configs';

// import App from './app';

// const app = new App();

// app.listen();




// process.env['NODE_CONFIG_DIR'] = __dirname + '/config';

// import dotenv from 'dotenv';
// import express from 'express';
// import bodyParser from 'body-parser';
// import { Subscriber } from './subscriber';
// import log4js from 'log4js';
// log4js.configure('./src/config/log4js.json');
// dotenv.config();

// class App {
//   public app: express.Application;
//   public port: string | number;
//   public env: string;
//   public log: log4js.Logger;

//   constructor() {
//     this.app = express();
//     this.port = process.env.PORT || 8080;
//     this.env = process.env.NODE_ENV || 'development';
//     this.log = log4js.getLogger('startup');

//     this.initializeMiddlewares();
//     this.initializeRoutes();
//   }

//   public listen() {
//     this.app.listen(this.port, async () => {
//       this.log.info(`App now running on port ${this.port} with pid ${process.pid}`);

//       try {
//         const subscriber = new Subscriber();
//         await subscriber.init();
//       } catch (e) {
//         this.log.error(e);
//         process.exit(1);
//       }
//     });
//   }

//   public getServer() {
//     return this.app;
//   }

//   private initializeMiddlewares() {
//     this.app.use(bodyParser.urlencoded({ extended: false }));
//     this.app.use(bodyParser.json());
//   }

//   private initializeRoutes() {
//     this.app.get('/', function (req, res) {
//       res.sendStatus(200);
//     });

//     this.app.post('/jobs', function (req, res) {
//       res.sendStatus(200);
//     });
//   }
// }

// export default App;
