// SPDX-License-Identifier: MIT 
pragma solidity >0.8.0;

import "./SolidityTest.sol";
// 远程导入（https://github.com/）

contract Hello {
    string public message = "Hello World!";
}
/**
 *学习参考文档： 
 * https://github.com/anbang/professional-solidity/blob/main/docs/source/02.type-of-data.md#1%EF%B8%8F%E2%83%A3-%E6%95%B0%E6%8D%AE%E4%B8%8E%E5%8F%98%E9%87%8F
 * @title
 * @author
 * @notice 向最终用户解释这是做什么的	
 * @dev 向开发人员解释任何额外的细节	
 * 合约地址/合约创建地址/合约调用地址
 *
 */
contract TypeOfData{

    address public owner;

    constructor() {
        // 可以用constructor 内获取当前合约地址
        owner=address(this);

         // 不可以在构造函数内调用函数，因为此时合约还没有完成构建好。
        // this.caller(); 相当于从外部调用 caller 方法
        // owner = this.caller();
    }

    function contractAds() external view returns (address){
        return address(this);
    }

    function name() external pure returns (string memory){
        return type(Hello).name;
    }
    function creationCode() external pure returns (bytes memory){
        return type(Hello).creationCode;
    }
     function runtimeCode() external pure returns (bytes memory) {
        return type(Hello).runtimeCode;
    }

    

    function test1() external pure returns(bool){
        uint256 a=100-99;

        uint256 b=100-1;

        if (a > 50 || b < 50) {
            return true;
        }
        return false;
    }
    // 计算gas 
     function testA2() external pure returns (bool) {
        if ((100 - 99) > 50 || (100 - 10) < 50) {
            return true;
        }
        return false;
    }
    int256 public minInt=type(int256).min;

    /**
    * int8 -2^7 ~ 2^7-1
    * uint8 0 ~ 2^8-1
    * 其他数据类型以此类推
    *
    * 问题: 为什么 uint8/int8至 uint256/uint256 都是以 8 的倍数递增，且最大值是 256?
    * 1 字节是 8 位，所以后面 8,16,都需要是 8 的整数倍，int8 是 8 位。
    * EVM 为地址设置的最大长度是 256 位，所以最大值是uint256/uint256。
    *
    * 问题: 字节 & bit & 十六进制数字关系?
    * bytes1 是指 1 个字节，1 个字节可以表示成 2 个连续的 16 进制数字。最大值是 0xff
    * bytes1 是指 1 个字节，1 个字节可以表示成 8 个连续的 bit 数字。最大值是 11111111
    * bytes1 等于两位连续的十六进制数字 0xXX
    * 8 个 bit 最大值是 11111111,8 个 bit 对应 2 个连续的十六进制数字，最大是 0xff;
    * uint8 等于两位连续的十六进制数字 0xXX 
    */
    function testA3() external pure returns (uint) {
        // max value is 2^256-1 uint 是无符号的，int 是有符号的
       uint256  u256Max = type(uint256).max;
       return u256Max;
    }
}
struct Book{
    string title;

    string author;

    uint256 book_id;
}

/**
 * 状态变量：存储在合约内部，相当于属于已经写入到区块链中
 * 局部变量：定义在特定函数中，仅仅在函数执行过程中生效
 * 全局变量：保存在全局命名空间，存在于EVM虚拟机中，不用定义，可以直接获取即可
        - msg.sender:表示当前调用方法时的发起人。合约的发起人可以是创建者，也可以是合约拥有者，**如何判断合约的拥有者？**

        - msg.value
        - block.timestamp
        - block.number
 *
 *   
 */
// 自由函数
function getBalance() view returns (uint256){
    return address(msg.sender).balance;
}
