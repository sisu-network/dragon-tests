const HDWalletProvider = require("@truffle/hdwallet-provider");
const ethers = require("ethers");
const fs = require("fs");
const assert = require("assert").strict;
require("dotenv").config();

const MNEMONIC = process.env.MNEMONIC;

const Web3 = require("web3");
let web3;
let provider;
let accounts;
let sendParams;

const sleep = function (ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
};

function getProvider() {
  const url = process.env.RPC_ENDPOINT;

  return new HDWalletProvider({
    mnemonic: {
      phrase: MNEMONIC,
    },
    providerOrUrl: url,
  });
}

async function init() {
  provider = getProvider();
  web3 = new Web3(provider);
  web3.setProvider(provider);

  accounts = await web3.eth.getAccounts();
  console.log(accounts);
  sendParams = {
    from: accounts[0],
    gasPrice: "75",
    gasLimit: "5000000",
  };
}

async function run() {
  await init();

  // const amountToSend = web3.utils.toWei("1", "ether");
  const amountToSend = "100000000000000000000"; // 100 eth

  for (let i = 1; i < accounts.length; i++) {
    console.log("Funding account", i);
    await web3.eth.sendTransaction({ from: accounts[0], to: accounts[i], value: amountToSend });
  }

  console.log("Done");
}

run();
