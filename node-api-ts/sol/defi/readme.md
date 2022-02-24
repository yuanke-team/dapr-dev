
# 去中心化的金融(Defi)
大家好，这次我们通过[ChainIDE](https://chainide.com/)在[Binance Smart Chain (BSC)](https://binanceide.com/project/welcome)上部署了**个简单的去中心化金融（Defi）智能合约项目。
这个项目包含两个基于ERC20的代币。我们暂且称这两个代币为*ChainToken*和*RevenueToken*（以下简称CT和RT）。CT用于模拟用户自己持有的数字货币，RT是用户认捐CT后获得的利息。整个Defi项目的工作流程由一个农场合同执行。用户将他/她持有的数字货币质押到农场合约中。该合同根据用户的质押金额生成RT，并将其作为利息返还给用户。
## 项目中使用的工具
1. [Binance IDE](https://binanceide.com/project/welcome)
2. [Metamask](https://metamask.io/) (连接到BSC Binance智能链)

## 连接方法
打开MetaMask小狐狸，选择自定义RPC选项。
<p align="center">
<img src="https://user-images.githubusercontent.com/16441258/103730398-bec83e80-501d-11eb-9655-c41a89607c45.png">
</p>

### 输入内容如下: 


Network Name: BSC Testnet


New RPC URL: https://data-seed-prebsc-1-s1.binance.org:8545/


Chain ID: 97


货币符号和区块资源管理器URL是可选的，可以留空。

<p align="center">
<img src="https://user-images.githubusercontent.com/16441258/103730883-d227d980-501e-11eb-8595-97d59134e94d.png">

## 操作程序
> 1. 点击[Binance IDE]主页上的Binance Smart Chain Docs案例（https://binanceide.com/project/welcome）。


<p align="center">
<img src="https://user-images.githubusercontent.com/16441258/103731380-f3d59080-501f-11eb-9a53-a3c05f256b96.png">
</p>

> 2. 在左边的目录中，我们准备了上述的三个智能合约。其中，ChainToken和Revenue Token是用于发行基于ERC20的代币的智能合约。合同中声明了数字货币的名称、发行总量、验证、转账和第三方转账。农场合同的代码实现了我们整个defi项目的操作逻辑。stakeToken方法是将CT抵押给这个合约，untakeTokens是取出抵押的数字货币，issueToken是根据抵押的CT产生RT利息。



<p align="center">
<img src="https://user-images.githubusercontent.com/16441258/103732087-a0fcd880-5021-11eb-949a-9ce6a67b10b1.png">
</p>

> 3.下面我们在Binance智能链上部署这三个合约，首先选择右上角的编译器，任何0.5.x的版本。


<p align="center">
<img src="https://user-images.githubusercontent.com/16441258/103732313-0f419b00-5022-11eb-8966-7f8b2b6f3682.png">
</p>

> 4. 点击编译后，你可以在输出控制台看到编译的结果。


<p align="center">
<img src="https://user-images.githubusercontent.com/16441258/103732477-5e87cb80-5022-11eb-8798-3e602c9ce4cd.png">
</p>
 
 

 
 > 5. 部署合同（设置值和wei）。在这一步中，你需要分别部署三个编译好的合约。 在编译合约中，首先部署ChainToken和RevenueToken，最后在部署Farm合约时传入前两个合约的地址。（前两个合约的地址可以在输出控制台查看，如 "contractAddress"。"0x79a377715E31D5F9eE736f8087aC0Ca230F8C48e")


<p align="center">
<img src="https://raw.githubusercontent.com/wkq1991zmc/defi/master/%E5%9B%BE%E7%89%87%E6%95%99%E7%A8%8B7.png">
</p>


<p align="center" >
<img src="https://raw.githubusercontent.com/wkq1991zmc/defi/master/%E5%9B%BE%E7%89%87%E6%95%99%E7%A8%8B8.png", width="500" height="700">
</p>

<p align="center" >
<img src="https://raw.githubusercontent.com/wkq1991zmc/defi/master/%E5%9B%BE%E7%89%87%E6%95%99%E7%A8%8B9.png", width="500" height="700">
</p>



> 6. 在所有合约部署完毕后，Farm合约实现了数字货币质押功能（stakeToken）。这个方法的实质是调用chainToken合约中的TransferFrom函数，将用户账户中的chainintoken转移到Farm合约中。我们需要在批准账户中拥有足够的数字货币。


<p align="center">
<img src="https://raw.githubusercontent.com/wkq1991zmc/defi/master/%E5%9B%BE%E7%89%87%E6%95%99%E7%A8%8B10.jpg">
</p>


> 7.批准完成后，进行认捐操作（填写你要认捐的金额）。


<p align="center">
<img src="https://raw.githubusercontent.com/wkq1991zmc/defi/master/%E5%9B%BE%E7%89%87%E6%95%99%E7%A8%8B11.png">
</p>

> 8. 接下来，RevenueToken被转移到Farm合约以产生红利，参数被传入Farm合约地址和需要转移的数字货币的数量。



<p align="center">
<img src="https://raw.githubusercontent.com/wkq1991zmc/defi/master/%E5%9B%BE%E7%89%87%E6%95%99%E7%A8%8B11.png", width="500" height="700">
</p>

<p align="center">
<img src="https://raw.githubusercontent.com/wkq1991zmc/defi/master/%E5%9B%BE%E7%89%87%E6%95%99%E7%A8%8B13.png", width="500" height="700">
</p>

> 9. 最后一步是调用农场合同中的issueToken函数，为已经认捐的用户产生红利。


<p><strong> 谢谢你阅读这篇文章。我们希望这篇文章能帮助你以简单的方式开发Defi应用程序。</strong></p>
 


