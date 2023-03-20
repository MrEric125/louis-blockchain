// SPDX-License-Identifier:MIT
pragma solidity ^0.8.13;

// Copyright (c) 2023 形参，
contract Purchase{
    uint public value;

    address public seller;

    address public buyer;

    enum State { Created, Locked, Inactive }
    State public state;

    constructor()  public payable{
        seller=msg.sender;
        value=msg.value/2;
        require((2*value)==msg.value,'value has to be even.');

    }
    modifier conditions(bool _condition){
        require(_condition);
        _;
    }
    modifier onlyBuyer(){
        require(msg.sender==buyer,'Only buyer can call this method');
        _;
    }
    modifier onlySeller() {
        require(
            msg.sender == seller,
            "Only seller can call this."
        );
        _;
    }
     modifier inState(State _state) {
        require(
            state == _state,
            "Invalid state."
        );
        _;
    }

    event Aborted();

    event PurchaseConfirmed();

    event ItemReceieved();

    function Abort() public onlySeller inState(State.Created){
        emit Aborted();

        state=State.Inactive;

        address(this).balance

        seller.transfer( );

    }

}