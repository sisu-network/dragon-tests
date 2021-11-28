const HDWalletProvider = require("@truffle/hdwallet-provider");
const ethers = require("ethers");
const fs = require("fs");
const assert = require("assert").strict;
require("dotenv").config();

const MNEMONIC = process.env.MNEMONIC;

const Web3 = require("web3");

const sleep = function (ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
};

function getProvider(url, mnemonic) {
  console.log("RPC_ENDPOINT = ", url);

  return new HDWalletProvider({
    mnemonic: {
      phrase: mnemonic,
    },
    providerOrUrl: url,
  });
}

async function getWeb3(type) {
  if (type == undefined || type == null) {
    type = "default";
  }

  let provider;
  switch (type) {
    case "ropsten":
      provider = getProvider(process.env.RPC_ROPSTEN, process.env.ROPSTEN_MNEMONIC);
      break;
    case "sisu_testnet":
      provider = getProvider(process.env.RPC_SISU_TESTNET, process.env.TESTNET_MNEMONIC);
      break;
    case "default":
      provider = getProvider(process.env.RPC_DEFAULT, process.env.MNEMONIC);
      break;
  }

  return new Web3(provider);
}

async function transfer() {
  const ropstenWeb3 = await getWeb3("ropsten");
  const testnetWeb3 = await getWeb3("sisu_testnet");

  // const web3s = [ropstenWeb3, testnetWeb3];
  // const amounts = ["15000000000000000" /** 0.015 eth */, "10000000000000000000" /** 10 eth */];
  // const recipients = [
  //   "0x341993A9d7B1792C92fE7A04Ec8053f3839Cf4F5",
  //   "0xBE3eFcee4d78c6489Cd1be4a8E6417f733B522BA",
  // ];

  const web3s = [testnetWeb3];
  const amounts = ["10000000000000000000" /** 10 eth */];
  const recipients = ["0xBE3eFcee4d78c6489Cd1be4a8E6417f733B522BA"];

  if (web3s.length != amounts.length || web3s.length != recipients.length) {
    throw new Error("Array lengths do not match");
  }

  for (var i = 0; i < web3s.length; i++) {
    let web3 = web3s[i];
    const accounts = await web3.eth.getAccounts();

    const balance = await web3.eth.getBalance(accounts[0]);
    console.log("accounts[0] = ", accounts);
    console.log("balance = ", balance);

    console.log("Transferring from", accounts[0], "to", recipients[i], ", value: ", amounts[i]);

    try {
      await web3.eth.sendTransaction({
        from: accounts[0],
        to: recipients[i],
        value: amounts[i],
      });
    } catch (error) {
      console.log(error);
    }

    console.log("Transfer done!");
  }
}

async function getBalance() {
  const testnetWeb3 = await getWeb3("ropsten");
  try {
    const accounts = await testnetWeb3.eth.getAccounts();
    console.log("accounts = ", accounts);

    const balance = await testnetWeb3.eth.getBalance("0xbeF23B2AC7857748fEA1f499BE8227c5fD07E70c");
    console.log("balance = ", balance);
  } catch (error) {
    console.log(error);
  }
}

// getBalance();
transfer();
