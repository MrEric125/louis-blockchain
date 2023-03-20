// SPDX-License-Identifier: UNLICENSED
// 表示不低于0.7.0的版本编译，但是不高于0.9.0
pragma solidity >=0.7.0 <0.9.0;


/// @title  委托投票
contract Ballot {

    struct Voter {

        uint weight;// 计票权重

        bool voted;// 若为真，代表该人已投票

        address delegate;// 被委托人

        uint  vote;// 投票天的索引
    }
    struct Proposal{
        bytes32 name;
        uint voteCount;
    }

    address public chairPerson;

    // todo mapping 的功能，状态变量的使用场景
    mapping (address=>Voter) public voters;

    Proposal[] public proposals;

    /**
    构造函数，创建一个临时的proposal 
     */
    constructor(bytes32[] memory proposalNames) public{
        chairPerson=msg.sender;
        voters[chairPerson].weight=1;
        for (uint256 i = 0; i < proposalNames.length; i++) {
            proposals.push(
                Proposal({
                    name:proposalNames[i],
                    voteCount:0
                    }));
            
        }
    }

    function giveRightToVote(address voter) public{
        // 若 require 的第一个参数的计算结果为false,则终止执行，撤销所有对状态和以太币余额的改动，类似java 中的checkCondition
        require(msg.sender==chairPerson,"only chairperson can give right to vote");

        require(!voters[voter].voted," the voter already voted");

        require(voters[voter].weight==0);

        voters[voter].weight==1;

    }

    function delegate(address to) public {
        Voter storage sender=voters[msg.sender];

        require(!sender.voted, "You already voted.");

        require(to != msg.sender, "Self-delegation is disallowed.");


        while(voters[to].delegate!=address(0)){

            to=voters[to].delegate;

        
            // 不允许闭环委托
            require(to != msg.sender, "Found loop in delegation.");
        }

        sender.voted=true;
        sender.delegate=to;
        Voter storage delegate_=voters[to];

        if(delegate_.voted){
            // 增加得票数
            proposals[delegate_.vote].voteCount+=sender.weight;
        }else{
            // 若被委托者还没投票，增加委托者的权重
            delegate_.weight+=sender.weight;
        }

    }

    function vote (uint proposal) public{
        Voter storage sender=voters[msg.sender];

        require(!sender.voted,"Already voted");

        sender.voted=true;

        sender.vote=proposal;
        // 如果 `proposal` 超过了数组的范围，则会自动抛出异常，并恢复所有的改动
        proposals[proposal].voteCount+=sender.weight;
    }

    function winningProposal() view public returns (uint winningProposal_) {
        uint winningVoteCount=0;

        for (uint256 index = 0; index < proposals.length; index++) {
            if(proposals[index].voteCount>winningVoteCount){
                winningVoteCount=proposals[index].voteCount;
                winningProposal_=index;
            }    
        }
        
    }

    function winnerName() view public returns (bytes32 winnerName_) {
        winnerName_=proposals[winningProposal()].name;
    }
}
