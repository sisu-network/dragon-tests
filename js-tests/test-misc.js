const HDWalletProvider = require("@truffle/hdwallet-provider");
const ethers = require("ethers");
const fs = require("fs");
const exec = require("child_process").exec;
const assert = require("assert").strict;

const MNEMONIC =
  "draft attract behave allow rib raise puzzle frost neck curtain gentle bless letter parrot hold century diet budget paper fetch hat vanish wonder maximum";
// const MNEMONIC = "ski oppose bird picture token burst drink pear owner turkey hair fiber";

const Web3 = require("web3");
let web3;
let provider;
let accounts;
let sendParams;

const sleep = function (ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
};

function getProvider() {
  const url = "HTTP://127.0.0.1:1234";

  return new HDWalletProvider({
    mnemonic: {
      phrase: MNEMONIC,
    },
    providerOrUrl: url,
  });
}

function getContractAbi(path) {
  let rawdata = fs.readFileSync(path);
  return JSON.parse(rawdata);
}

async function deployContract() {
  const contractAbi = getContractAbi("../go-tests/tests/contracts/erc20/ERC20.abi");
  const contractByteCode = fs.readFileSync("../go-tests/tests/contracts/erc20/ERC20.bin", "utf8");
  const contract = new web3.eth.Contract(contractAbi);

  try {
    return await contract
      .deploy({
        data: contractByteCode,
        arguments: ["Test", "$TTTT"],
      })
      .send(sendParams);
  } catch (err) {
    console.log(err);
    return null;
  }
}

async function testWeb3Call() {
  const contract = await deployContract();

  console.log("Getting balance...");

  try {
    await contract.methods.mint(accounts[0], 1).send(sendParams);
    balance = await contract.methods.balanceOf(accounts[0]).call();
    console.log("Account[0] first balance = ", balance);

    await contract.methods.transfer(accounts[1], 1).send(sendParams);

    balance = await contract.methods.balanceOf(accounts[0]).call();
    console.log("Account[0] balance after transfer = ", balance);

    balance = await contract.methods.balanceOf(accounts[1]).call();
    console.log("Account[1] balance after transfer = ", balance);
  } catch (err) {}
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
  await testWeb3Call();
}

run();
