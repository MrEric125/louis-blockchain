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
