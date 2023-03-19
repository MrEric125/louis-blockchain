
// 测试脚本

import { time, loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs.js";
import { expect } from "chai";

const {ethers }=require("hardhat");

describe("Token", function () {
  let Token ,token ,owner,address,secret 
  beforeEach(async function () {
    Token=await ethers.getContractFactory("BigToken");
    token =await Token.deploy;
    [owner,address,secret]=await ethers.getSigners();

  });
  describe("Token", function () {
    if("success owner",async () =>{
      expect( await token.owner()).to.equal(owner.address);
    });
  });
})
