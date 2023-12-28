// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.19;

/* payable: 
*       - 可标记函数: 可以执行接收以太的这种交易
*       - 标记地址：  标记地址表示可以向这个地址转货币
* fallback(回退函数):  todo(深入理解使用场景)
    - fallback() external [payable]
    - fallback(types calldata input ) external [payable] returns (bytes memory output)
    - 在两种情况下会被调用，
        - 向合约转账
        - 执行合约不存在的方法

*/
contract PayableTest{

    event Log(string funName, address from,uint value,bytes data);

    function deposit() external payable {}

    function withRaw() external {
        payable(msg.sender).transfer(address(this).balance);
    }

     function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
    fallback() external payable { 
        emit Log("fallback", msg.sender, msg.value, msg.data);
    }

    
}