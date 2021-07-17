const express = require('express');
const { exec, execSync } = require('child_process');
const { Validator } = require('@chainlink/external-adapter');

const app = express();
const port = process.env.EA_PORT || 8080;

app.post('/', (req, res) => {
  const validator = new Validator(req.body, {});

  execSync('bootnode --genkey ./bootnode.key');
  const bootnode = exec('bootnode --nodekey ./bootnode.key --addr :5550 --group 0 --nodes 3', (error, stdout, stderr) => {
  });
  bootnode.stdout.on("data", function (data) {
    const startPos = data.indexOf("enode");
    if (startPos === -1)
      return;
    const endPos = data.indexOf("@");
    res.send({ jobRunID: validator.validated.id, data: { enode: data.slice(startPos, endPos) } });
  });
});

app.listen(port, () => { });