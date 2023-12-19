// SPDX-License-Identifier: MIT 
pragma solidity >=0.8.0;

/**
* 
*
*
*/
contract testArray{

    enum FreshJuiceSize {SMALL,MEDIUM,LARGE}

    FreshJuiceSize choice;

    FreshJuiceSize constant defaultChoice=FreshJuiceSize.MEDIUM;

    function setLarge() public {
        choice=FreshJuiceSize.LARGE;
    }
    function getChoice() public view returns(FreshJuiceSize){
        return choice;
    }
    function getDefaultChoice() public pure returns (uint) {
      return uint(defaultChoice);
   }

   mapping(address=> uint) public balances;

   function updateBalance(uint newBalance) public {
     balances[msg.sender]=newBalance;

   }

    function testArrays() public pure {
        uint len =7;

        uint [] memory a=new uint[](7);

        bytes memory b=new bytes(len);

        assert(a.length==7);

        assert(b.length==len);

        a[6]=8;

        assert(a[6]==8);

        uint[3] memory c=[uint(1),2,3];

        assert(c.length==3);
    }
}
contract Updater {
   function updateBalance() public returns (uint) {
      testArray ledgerBalance = new testArray();
      ledgerBalance.updateBalance(10);
      return ledgerBalance.balances(address(this));
   }
}
