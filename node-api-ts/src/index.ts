import { DaprServer, HttpMethod, CommunicationProtocolEnum } from "dapr-client";
import Web3  from "web3";
const Tx = require('ethereumjs-tx').Transaction
const common = require('ethereumjs-common');
// import fs  from "fs";

const daprHost = "127.0.0.1";
const daprPort = "50050"; // Dapr Sidecar Port of this Example Server
const serverHost = "127.0.0.1"; // App Host of this Example Server
const serverPort = "50051"; // App Port of this Example Server

// mainnet 
// const web3Net = new Web3('https://bsc-dataseed1.binance.org:443');

// testnet
const web3TestNet = 'https://data-seed-prebsc-1-s1.binance.org:8545';

async function start() {
  const server = new DaprServer(serverHost, serverPort, daprHost, daprPort, CommunicationProtocolEnum.GRPC);
  await server.startServer();
  // web3 初始化
  var web3 = new Web3(web3TestNet)

  await server.invoker.listen("hello-world", async (data: any) => {
    console.log("[Dapr-JS][Example] Received Hello World Method Call");
    console.log(`[Dapr-JS][Example] Data: ${JSON.stringify(data.body)}`);
    console.log(`[Dapr-JS][Example] Replying to the client`);
    return { hello: "world received" };
  }, { method: HttpMethod.POST });

  // 币安 hello-world demo 调用
  await server.invoker.listen("bsc-hello-world", async (data: any) => {
    console.log(`[Dapr-BSC] Data: ${JSON.stringify(data.body)}`);

    // 合约ABI
    var abi = JSON.parse('[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"say","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]');
    // var abi=JSON.parse(fs.readFileSync("/home/liaoxin/dev/dapr/dapr-dev/node-api-ts/sol/hello_sol_hello.abi").toString())

    // 合约地址
    var address = "0xbc961C8A39E93156C1AD75858BC281a24eeFd266";
    // 通过ABI和地址获取已部署的合约对象
    var helloContract = new web3.eth.Contract(abi,address);

    // 调用智能合约方法
    var helloResult
    await helloContract.methods.say().call().then(function(result :any){
      console.log('show the custom say call:');
      console.log(result)
      helloResult = result
    })
    return { hello: helloResult };
  }, { method: HttpMethod.POST });

  // 币安 Storage 调用
  await server.invoker.listen("bsc-storage", async (data: any) => {
    console.log(`[Dapr-BSC][bsc-storage] Data: ${JSON.stringify(data.body)}`);

    
    if (data.body.length == 0){
      console.log("body is empty");
    }else{
     var body = JSON.parse(data.body);
    }

    // console.log(body.num)
    // 合约ABI
    var abi = JSON.parse('[{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"num","type":"uint256"}],"name":"store","outputs":[],"stateMutability":"nonpayable","type":"function"}]');
    
    // 合约地址
    var contractAddress = "0x846828c7F7603a9e75Cd33Af879A1dE6BbdF5A69";


    // 私钥转换为Buffer
    const privateKey =  Buffer.from('b9fd52f015005958799421d06e2476f2b5ab7831a257fd6df3c53422fe0b5eae',"hex")
    // //私钥转换为账号
    const account = web3.eth.accounts.privateKeyToAccount('0xb9fd52f015005958799421d06e2476f2b5ab7831a257fd6df3c53422fe0b5eae');
    
    // //私钥对应的账号地地址
    const address = account.address
    console.log("user address: ",address)

    // 通过ABI和地址获取已部署的合约对象
    var contract = new web3.eth.Contract(abi,contractAddress);

    //调用合约abi store（） 方法
    var contractData =  contract.methods.store(body.num).encodeABI() 
              
    // 调用智能合约返回结果
    var helloResult

    if (data.body.length > 0 ){
      if (body.num == undefined){
        return { message : "num is null"}
      }
      web3.eth.getChainId((err, chainid) => {
        if (err != null) {
          console.log('chainid:' + chainid + ' err:' + err);
          return;
        }
        console.log('chainid:' + chainid);
        // config custom chain: BSC testnet
        const chain = common.default.forCustomChain(
          'mainnet',{
              name: 'bnbt',
              networkId: chainid,
              chainId: chainid
          },
          'petersburg'
        );

        //获取nonce,使用本地私钥发送交易
        web3.eth.getTransactionCount(address).then(
          nonce => {
          
              console.log("nonce: ",nonce)

              const txParams = {
                  nonce: web3.utils.toHex(nonce),
                  gasLimit : web3.utils.toHex(30000),  //408587 15% extra 
                  gasPrice : web3.utils.toHex(web3.utils.toWei('10','gwei')), //web3.utils.toHex(6000000000), 20% extra 
                  to: contractAddress,
                  // from: address,
                  // value: 0,//web3.utils.toHex(web3.utils.toWei('0.001', 'ether')),  //0.01
                  data: contractData,
              }

              console.log(`[Dapr-BSC][bsc-storage] txParams: ${JSON.stringify(txParams)}`)
              const tx = new Tx(txParams, {common: chain})
              tx.sign(privateKey)
              const serializedTx = tx.serialize()
              const row = "0x" + serializedTx.toString('hex')
              
              web3.eth.sendSignedTransaction(row, (err, txHash) => {
                var txHash = 'https://testnet.bscscan.com/tx/'+txHash
                console.log('txHash ', txHash )
              }).on('error', function(error) {
                console.log(error);
              });
          },
          e => console.log(e)
        )
      });
      return { store: "sany success" };
    }else{
      await contract.methods.retrieve().call().then(function(result :any){
        console.log('show the custom retrieve call:', result);
        helloResult = result
      });
      return { retrieve: helloResult };
    }
  }, { method: HttpMethod.POST });

}

start().catch((e) => {
    console.error(e);
    process.exit(1);
});