pragma solidity ^0.8.0;

contract HelloWorld {

    string public message ="HelloWorld";// 状态变量

    function fn1() public view returns(string memory){
        return message;
    }

    function fn2() public pure returns(string memory){
        return "hello world add in remix";
    }

    uint256 storeData;

    /// 这种注释会比较有意义，会生成文档 netspect 注释
    /// @param _x
    function set(uint256 _x)  public{
        storeData=_x;
    }


    // this 本身代表当前合约
    // address 当前合约地址
    function fn1() external view returns(address){
        return address(this); // 当前合约地址
    }

    address public owner;
    address public owner2;


    constructor(){
        owner=address(this); // 合约地址，
        owner2=msg.sender;

    }
    function fn3() public  view returns(address){
        return msg.sender;
    }


    // 合约结构接受
    // 1. spdx 版本声明
    // 2。 pragma solidity 版本限制;
            // 变量，函数，this 关键字，合约地址，合约创建者地址，合约调用地址，合约属性type关键字
    // 3. import 声明
    // 4. interface 接口
    // library 库合约


    // function fn3() public pure return(string memory){
    //     return fn2();
    // }

}
