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
    uint public storedData = 30;
    uint public data = 30;

    // 状态变量的三种作用域: public Internal private
    constructor() {

    }
    function getResult() public pure returns (string memory){

        uint a = 1; // local variable
        uint b = 2;
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


}

// contract Caller {
//     SolidityTest c = new SolidityTest();

//     function f() public view returns (uint) {
//         return c.storedData(); //external access
//     }
// }