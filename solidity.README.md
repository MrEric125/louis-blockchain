### 本地连接remix 

官方文档： https://remix-ide.readthedocs.io/en/latest/remixd.html


1. 安装 
> sudo npm install -g @remix-project/remixd
> remixd -s ./ --remix-ide https://remix.ethereum.org
3. 直接访问 https://remix.ethereum.org 即可，我的版本是0.6.8

国内安装可先切换镜像源到淘宝镜像
> npm install -g cnpm --registry=https://registry.npmmirror.com
> npm config set registry https://registry.npmmirror.com

window11 可能会系统会提示
```
无法加载文件 C:\Users\louis\AppData\Roaming\npm\remixd.ps1，因为在此系统上禁止运行脚本。
```

需要用管理员身份运行 powerShell 

执行 get-ExecutionPolicy 会形式可执行的状态

> Restricted: 禁止的
> RemoteSigned  允许

执行：set-ExecutionPolicy 输入RemoteSigned 即可

### smart contract 

 账户，私钥，公钥，地址，密码，助记词，钱包，

 gas 费用计算

### tool 
Dapp: 
solidity REPL
solgraph
``
## 区块链基础
### 交易、事务
### 区块
## 以太坊虚拟机
### 账户，交易，gas,存储，内存和栈，指令集，消息调用，委托调用，代码调用和相关库，日志
------
语法文档: 
https://www.tutorialspoint.com/solidity/solidity_overview.htm
您可以使用 Solidity 创建用于投票、众筹、盲拍和多重签名钱包等用途的合约。

## 安装
> npm install -g solc

### 变量名称规则
1. 不应使用保留字作为变量名称
2. 不能以0-9 这些数字开头，必须以字母或下滑线开头
3. 区分大小写


### Global Variables
These are special variables which exist in global workspace and provide information about the blockchain and transaction properties.
| Name | Returns |
|---|---|
| blockhash(uint blockNumber) returns (bytes32) | Hash of the given block - only works for 256 most recent, excluding current, blocks |
| block.coinbase (address payable) | Current block miner's address |
| block.difficulty (uint) | Current block difficulty |
| block.gaslimit (uint) | Current block gaslimit |
| block.number (uint) | Current block number |
| block.timestamp (uint) | Current block timestamp as seconds since unix epoch |
| gasleft() returns (uint256) | Remaining gas |
| msg.data (bytes calldata) | Complete calldata |
| msg.sender (address payable) | Sender of the message (current caller) |
| msg.sig (bytes4) | First four bytes of the calldata (function identifier) |
| msg.value (uint) | Number of wei sent with the message |
| now (uint) | Current block timestamp |
| tx.gasprice (uint) | Gas price of the transaction |
| tx.origin (address payable) | Sender of the transaction |
