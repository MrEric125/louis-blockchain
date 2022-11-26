pragma solidity ^0.8.0;

contract HelloWorld {

    string public message ="HelloWorld";// 状态变量

    function fn1() public view returns(string memory){
        return message;
    }

    function fn2() public pure returns(string memory){
        return "hello world";
    }
    function fn3() publicpure return(string memory){
        return fn2();
    }


    function HelloWorld(){

    }
}
