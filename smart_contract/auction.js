
const web3=require('web3');

const provider=new web3.providers.HttpProvider('http://localhost:8545');

web3.setProvider(provider);

const contract=new web3.eth.Contract(abi,address);

const web3 = new Web3(rpcURL)

const address = "0x03118E2c88676d31ee397E1eEf7789fECfbC40b9"

web3.eth.getBalance(address,(err,wei)=>{

  console.log(wei);
  web3.utils.fromWe(wei,'ether')
  console.log("balance: "+balance)
});

var Tx =require('ethereumjs-tx').Transaction;

const Web3 = require('web3');


const account1='';

const privateKey1=Buffer.from("private key 1",'hex');

const txObj={
    nonce: web3.utils.toHex(txCount),
    gasPrice:web3.utils.toHex(100000),
    gasLimit:0,
    to:address,
    data:data,
    value:web3.utils.toWei("1","ether"),
};

const contractABI=[
    {
        "constant": false,
        "inputs": [
            {
                "name": "_to",
                "type": "address",
                "internalType":"string",
    }
]