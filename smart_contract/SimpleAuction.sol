//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

import "hardhat/cosole.sol";

contract SimpleAuction{
    address public beneficiary;

    uint public auctionEnd;

    // uint public now;


    address public highestBidder;

    uint public highestBid;

     mapping ( address =>uint)  pendingReturns;

    // 拍卖结束后设为true,禁止所有的变更
    bool ended;

    event HighestBidIncreased(address bidder,uint amount);


    event AuctionEnded(address winner,uint amount);

    constructor(uint _biddingTime, address _beneficiary) public{
        beneficiary=_beneficiary;
        auctionEnd= block.timestamp+_biddingTime;

    }

    function bid() public{
        // 如果拍卖已结束，撤销函数的调用。
        require(
             block.timestamp <= auctionEnd,
            "Auction already ended."
        );

        // 如果出价不够高，返还你的钱
        require(
            msg.value > highestBid,
            "There already is a higher bid."
        );
        if (highestBid != 0) {
            // 返还出价时，简单地直接调用 highestBidder.send(highestBid) 函数，
            // 是有安全风险的，因为它有可能执行一个非信任合约。
            // 更为安全的做法是让接收方自己提取金钱。
            pendingReturns[highestBidder] += highestBid;
        }
        highestBidder = msg.sender;
        highestBid = msg.value;
        emit HighestBidIncreased(msg.sender, msg.value);
        
    }
    function withdraw() public returns(bool){
        uint amount = pendingReturns[msg.sender];

        if(amount>0){
            pendingReturns[msg.sender] =0;
            if(!msg.sender.send(amount)) {
                pendingReturns[msg.sender] =amount;
                return false;


            }
        }
        return true;
    }

    function auctionEnd_() public {
        require(now>=auctionEnd,"auction not ye ended.");
         require(!ended, "auctionEnd has already been called.");
         ended = true;
         emit AuctionEnded(highestBidder, highestBid);

         beneficiary.transer(highestBid);
    }
}
