// SPDX-License-Identifier: MIT 
pragma solidity >=0.5.0 <0.10.0;

// 支持使用双引号("")，和单引号('')的字符串
contract SolidityTest{
    string data="test";

    string data2='test2';

    /* string type requires more gas,
     * Solidity 提供字节到字符串之间的内置转换，反之亦然。
     * 在 Solidity 中，我们可以轻松地将字符串文字分配给 byte32 类型变量 
     */
    bytes32 data3="test3";


    function c(uint _i) public pure returns (string memory) {
        if (_i == 0) {
         return "0";
        }
        uint len;
      
        uint j = _i;
        while (j != 0) {
            len++;
            j /= 10;
        }
        /* 
         *
         * Note: The called function should be payable if you send value and 
         * the value you send should be less than your current balance.
         * Debug the transaction to get more information.
         *
         */
        bytes memory bstr = new bytes(len);
        uint k = len - 1;
      
        while (_i != 0) {
            bstr[k--] = bytes1(uint8(48 + _i % 10));
            _i /= 10;
        }
      return string(bstr);

    }

   
}