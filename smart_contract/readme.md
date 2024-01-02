## function 
### 到底什么是 Pure 和View？

刚开始学习 solidity 时，`pure` 和 `view` 关键字可能令人费解，因为其他编程语言中没有类似的关键字。solidity 引入这两个关键字主要是因为 以太坊交易需要支付气费（gas fee）。合约的状态变量存储在链上，gas fee 很贵，如果计算不改变链上状态，就可以不用付 gas。包含 pure 和 view 关键字的函数是不改写链上状态的，因此用户直接调用它们是不需要付 gas 的（注意，合约中非 `pure/view` 函数调用 `pure/view` 函数时需要付gas）。

在以太坊中，以下语句被视为修改链上状态：

1. 写入状态变量。
2. 释放事件。
3. 创建其他合约。
4. 使用 `selfdestruct`.
5. 通过调用发送以太币。 
6. 调用任何未标记 `view` 或 `pure` 的函数。
7. 使用低级调用（low-level calls）。
8. 使用包含某些操作码的内联汇编。

* `view` 函数能读取但也不能写入状态变量
* `pure` 函数既不能读取也不能写入链上的状态变量

## dataStorage

Solidity数据存储位置有三类：`storage`，`memory`和`calldata`。不同存储位置的gas成本不同。`storage`类型的数据存在链上，类似计算机的硬盘，消耗gas多；`memory`和`calldata`类型的临时存在内存里，消耗gas少。大致用法：

1. `storage`：合约里的状态变量默认都是storage，存储在链上。

2. `memory`：函数里的参数和临时变量一般用memory，存储在内存中，不上链。

3. `calldata`：和`memory`类似，存储在内存中，不上链。与`memory`的不同点在于`calldata`变量不能修改（immutable），一般用于函数的参数。

## mapping 

在映射中，人们可以通过键（Key）来查询对应的值（Value），比如：通过一个人的id来查询他的钱包地址。

声明映射的格式为`mapping(_KeyType => _ValueType)`，其中`_KeyType`和`_ValueType`分别是Key和Value的变量类型。例子：
```solidity

mapping(uint => address) public idToAddress; // id映射到地址
mapping(address => address) public swapPair; // 币对的映射，地址到地址
```
### 映射的规则

* 规则1：映射的`_KeyType`只能选择Solidity`内置的值类型`，比如uint，address等，不能用自定义的结构体。而`_ValueType可以使用自定义的类型`。下面这个例子会报错，因为_KeyType使用了我们自定义的结构体：
```solidity

// 我们定义一个结构体 Struct
struct Student{
    uint256 id;
    uint256 score;
}
mapping(Student => uint) public testVar;
```
* 规则2：映射的存储位置必须是`storage`，因此可以用于合约的状态变量，函数中的storage变量和library函数的参数。不能用于public函数的参数或返回结果中，因为mapping记录的是一种关系 (key - value pair)。

* 规则3：如果映射声明为public，那么Solidity会自动给你`创建一个getter函数`，可以通过Key来查询对应的Value。

* 规则4：给映射新增的键值对的语法为`_Var[_Key] = _Value`，其中_Var是映射变量名，_Key和_Value对应新增的键值对。例子：
```solidity

function writeMap (uint _Key, address _Value) public{
    idToAddress[_Key] = _Value;
}
```
### 映射的原理
* 原理1: 映射不储存任何键（Key）的资讯，也没有length的资讯。

* 原理2: 映射使用`keccak256(abi.encodePacked(key, slot))`当成offset存取value，其中slot是映射变量定义所在的插槽位置。

* 原理3: 因为Ethereum会定义所有未使用的空间为0，所以未赋值（Value）的键（Key）初始值都是各个type的默认值，如uint的默认值是0。
