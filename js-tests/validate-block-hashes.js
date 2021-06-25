const axios = require("axios");
require("dotenv").config();

function decimalToHexString(number) {
  if (number < 0) {
    number = 0xffffffff + number + 1;
  }

  return number.toString(16).toUpperCase();
}

async function getData(ip, blockNumber) {
  axios.defaults.baseURL = `http://${ip}:1234`;

  const headers = {
    "Content-Type": "application/json",
  };

  const response = await axios.post(
    "/",
    { jsonrpc: "2.0", method: "eth_getBlockByNumber", params: [blockNumber, false], id: 1 },
    headers
  );

  return response.data;
}

async function check() {
  const ips = process.env.NODE_IPS.split(",");

  for (let k = 2000; k <= 2010; k++) {
    const blockNumber = "0x" + decimalToHexString(k);

    console.log("k = ", k);

    const data = [];
    for (let i = 0; i < ips.length; i++) {
      const result = await getData(ips[i], blockNumber);
      data.push(result);
    }

    for (let i = 1; i < ips.length; i++) {
      if (data[0].result.hash != data[i].result.hash) {
        for (let j = 0; j < ips.length; j++) {
          console.log(`data ${j} = ${data[j]}`);
        }

        throw Error("Data hashes are not the same.");
      }
    }

    console.log("Data is correct");
  }

  // const latest = await getData("44.193.169.123", "latest");
  // // const blockNumber = latest.result.number;
}

async function run() {
  await check();
}

run();
