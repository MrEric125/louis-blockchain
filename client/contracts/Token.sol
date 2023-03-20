

contract Token{

    string public name ='my_token';

    string public symbol ='NTN';

    uint256 public totalSupply=10000;

    address public owner;

    mapping (address =>uint) balances;

    constructor(){
        owner=msg.sender;
        balances[msg.sender] = totalSupply;
    }

    function transfer(address to ,uint amount) external {
        balances[msg.sender]-=amount;
        balances[to]+=amount;

    }
    function balanceOf(address acount) external view  returns(uint) {
        return balances[acount];
    }
}