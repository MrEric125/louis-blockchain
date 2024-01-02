// SPDX-License-Identifier: MIT 
pragma solidity <0.9.0;

// UniswapV2 闪电贷回调接口
interface IUniswapV2Callee{
    function uniswapV2Call(address sender,uint amount0,uint amount1,bytes calldata data) external;
}

contract UniswapV2Flashload is IUniswapV2Callee{

    address private constant UNISWAP_V2_FACTORY="0x123";

    

    function uniswapV2Call(address sender,uint amount0,uint amount1,bytes calldata data) external{

    }

}