// SPDX-License-Identifier: MIT 
// 无法编译早于 0.8.10 版本开始的编译器上运行的文件的编译指示将编写如下
// ^ 0.8.10  则是只允许 0.8.x 对该代码进行编译，不允许 0.7.x 或者0.9.x 的编译器编译合约
pragma solidity >=0.8.10;

contract SimpleStorage{
    uint storedData;

    function set(uint x) public{
        storedData=x;
    }

    function get () public view returns(uint){
        return storedData;
    }
}