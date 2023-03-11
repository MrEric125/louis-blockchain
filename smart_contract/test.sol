
pragma solidity ^0.8.0;


contract test{

    uint256 a=10 ;

    uint256 b;

    uint256 c;

    string name ="eric";

    function getA() public view returns(uint256){
        return a;
    }

   /*
   * view 查看链上数据，不改变数据值
   * pure 纯粹的查看函数返回，数据不上链
   */
    function pretest(uint256 param) public pure returns(uint256){
         return param;
    }
}