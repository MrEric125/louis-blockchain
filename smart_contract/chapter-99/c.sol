// SPDX-License-Identifier: MIT 
pragma solidity >=0.5.0 <0.9.0;


/**
 * if ,else,while,do,for,break,continue,return,都可以使用，
 * try /catch 也可以使用，但是只适用于外部函数调用，和合约创建调用，可以使用恢复状态来创建错误。
 */
contract c {
    function g(uint a) public pure returns(uint ret){ return a+f();}

    function f() internal pure returns (uint ret){ return g(7)+f();}

    /* 
     * 加盐合约创建
     */
    function createDSalted(bytes32 salt,uint arg) public {
        address predictedAddress=address(uint160(uint(keccak256(abi.encodePacked(
            bytes1(0xff),
            address(this),
            salt,
            keccak256(abi.encodePacked(type(D).creationCode,abi.encode(arg)
            ))
        )))));

        D d = new D{salt: salt}(arg);
        require(address(d) == predictedAddress);
        
    }
}

contract D{
    uint public x;

    constructor(uint a){
        x=a;
    }
}