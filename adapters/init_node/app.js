const express = require('express');
const { exec, execSync } = require('child_process');
const { Validator } = require('@chainlink/external-adapter');

const app = express();
const port = process.env.EA_PORT || 8080;

const customParams = {
  enode: ['enode'],
  ip: ['ip']
}

app.post('/', (req, res) => {
  const validator = new Validator(req.body, customParams);
  const data = req.body.data;
  execSync('gdcrm --genkey node.key');
  exec(`gdcrm --rpcport 9011 --bootnodes '${data.enode}@${data.ip}:5550' --port 12341 --nodekey 'node.key'`, (error, stdout, stderr) => {

  });
  res.status(200).json({ jobRunID: validator.validated.id, data: { } });
});

app.listen(port, () => { });