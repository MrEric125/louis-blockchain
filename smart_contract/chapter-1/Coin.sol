// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.19;

contract Coin{
    uint storageData;

    function set(uint x) public {
        storageData=x;
        
    }

    /**
        声明可以被公开访问address 类型的状态变量
        类型+访问权限+变量名
        address：是一个160位的值，不允许任何算数操作，相当于是一个常量，适合存储合约地址或外部人员的秘钥对
        public：是访问权限，含义同其他语言中的public，这个变量除了被合约访问以外，其他合约也可访问

        编译器会把这一样代码编译成，
        function minter() returns(address){return minter;}
     */
    address public minter;

    /* 
        创建一个 公共类型，
        将address映射为无符号整数
        mapping 可以看成是一个哈希表，会执行虚拟初始化

     */
    mapping(address=> uint) public balance;


    event Send(address from,address to,uint amount);

    // constructor Coin() public{
    //     minter=msg.sender;
    // }
}

