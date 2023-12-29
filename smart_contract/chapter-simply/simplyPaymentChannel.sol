// SPDX-License-Identifier: MIT 
pragma solidity >=0.8.0;

contract SimplyPaymentChannel{
    address payable public sender;

    address payable public recipient;

    uint256 public expiration;

    address owner=msg.sender;

    mapping(uint256=> bool) usedNonces;



    constructor(address payable recipientAddress,uint256 duration) public payable{
        sender=payable(msg.sender);
        
        recipient=recipientAddress;

        expiration=block.timestamp+duration;

    }

    function close(uint256 amount,bytes memory signature) external{
        require(msg.sender==recipient);

        require(isValidSignature(amount, signature));

        recipient.transfer(amount);

        selfdestruct(sender);
    }
    function extend(uint256 newExpiration) external {
        require(msg.sender==sender);

        require(newExpiration>expiration);

        expiration=newExpiration;
    }

    function claimPayment(uint256 amount,uint256 nonce,bytes memory signature) external{
        require(!usedNonces[nonce]);

        usedNonces[nonce]=true;
        //  重建客户端签名信息
        bytes32 message=prefixed(keccak256(abi.encodePacked(msg.sender,amount,nonce,this)));

        require(recoverSigner(message, signature)==owner);

        payable(msg.sender).transfer(amount);


    }


    function isValidSignature(uint256 amount,bytes memory signature) internal view returns(bool){
        bytes32 message=prefixed(keccak256(abi.encodePacked(this,amount)));

        return recoverSigner(message,signature)==sender;
    }

    function prefixed(bytes32 hash) internal pure returns(bytes32){
        return keccak256(abi.encodePacked("\x19Etherem signed Message:\n32",hash));
    }
    
    function recoverSigner(bytes32 message,bytes memory sig) internal pure returns(address){
        (uint8 v, bytes32 r, bytes32 s)=splitsSignature(sig);

        return ecrecover(message,v,r,s);
    }
    // function constructPaymentMessage(contractAddress,amount) public{
    //     return abi.soliditySHA3(['address','uint256'],[contractAddress,amount]);
    // }

    // function signMessage(message,callback) public{
    //     web3.eth.personal.sign("0x"+message.toString("hex"),web3.eth.defaultAccount,callback);
    // }

    /**
     * 分离签名信息
     */
    function splitsSignature(bytes memory sig) internal pure returns(uint8 v,bytes32 r, bytes32 s){
        require(sig.length==65);

        assembly{
            r:= mload(add(sig,32))

            s:= mload(add(sig,64))

            v:= byte(0,mload(add(sig,96)))
        }
        return (v,r,s);
    }


}