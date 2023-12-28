// SPDX-License-Identifier: MIT 
pragma solidity >=0.8.0;

// 很简单的合约
/* 
 * solidity supports three types of variables
    1. state variables: 其值永久存储在合约存储中的变量
    2. local variables: 其值在函数执行之前一直存在的变量。
    3. global variables: 全局命名空间中存在特殊变量，用于获取有关区块链的信息。
 *
 */
contract SolidityTest {

    // state variables
    // uint public storedData = 30;
    // uint public data = 30;

    // 状态变量的三种作用域: public Internal private
    // constructor() {

    // }

    /**
    * https://github.com/etherchina/solidity-doc-cn/blob/defb2004c8c5c080ab147af94b4d95334bbb6002/miscellaneous.rst#L347
    * function: the basic syntax is shown here
    * function function_name(parameter-list) scope returns ()
    * scope 种类:
            view:
            pure: 
            constant:
            payable:
            anonymous:
            indexed:
    */
    function getResult() public pure returns (string memory){

        uint a = 1; // local variable
        uint b = 0;
        uint result = a + b;
        return intToString(result); //access the state variable
    }

    function intToString(uint _i) internal pure returns (string memory){
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len - 1;

        while (_i != 0) { // while loop
            bstr[k--] = bytes1(uint8(48 + _i % 10));
            _i /= 10;
        }

        return string(bstr);

    }

    /**
    * The transaction has been reverted to the initial state.
    * Note: The called function should be payable if you send value and the value you send should be less than your current balance.
    * Debug the transaction to get more information.
    */
    function getResult2() public pure  returns (string memory){

        uint a = 1; // local variable
        uint b = 2;
        uint _i = a + b;

        uint j=0;
        uint len;
        for (j = _i; j != 0; j /= 10) {  //for loop example
            len++;         
        }
        bytes memory bstr = new bytes(len);
        uint k = len - 1;
        while (_i != 0) {
            bstr[k--] = bytes1(uint8(48 + _i % 10));
            _i /= 10;
        }
        return string(bstr);//access local variable

        // return intToString(result); //access the state variable
    }

    function integerToString2(uint _i) internal pure 
      returns (string memory) {
      if (_i == 0) {
         return "0";
      }
      uint j=0;
      uint len;
      for (j = _i; j != 0; j /= 10) {  //for loop example
         len++;         
      }
      bytes memory bstr = new bytes(len);
    //   uint k = len - 1;
    //   while (_i != 0) {
    //      bstr[k--] = bytes1(uint8(48 + _i % 10));
    //      _i /= 10;
    //   }
      return string(bstr);//access local variable
   }
}
