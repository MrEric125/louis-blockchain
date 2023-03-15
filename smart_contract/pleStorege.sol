// SPDX-License-Identifier:MIT
pragma solidity ^0.8.13;


// import 是导入外部合约，
// import ".abi";




/**
 合约基础语法：对象创建， 继承，实现， 
 容器存取，遍历，
 文件读取，存放
 


 */

/**
**
*
*/
contract SimpleStorage{

 uint favoriterNum;

 

 struct People{
     uint age;

     string name;
 }

 People public person1=People(30,"Fredhan");


 People[] public persons;

 mapping (string =>uint) public NameToAge;



 function setFavoriterNum(uint _fterNum) public{
     favoriterNum=_fterNum;

 } 
 function retrieve() public view returns(uint){
     return favirateNum;
 }  
 // memory and calldate 
 // calldate 形参内容不能被修改，memory 形参内容能被修改
 function addPerson(uint _age,string memory _name) public{
     persons.push(People(_age,_name));
     NameToAge[_name]=_age;
 }

 // inject provider 
 function testNetwork() public returns(uint){
    
 }

 function store(uint _num) public virtual   {

 }
}

contract StorageFactory{
    SimpleStorage[] public simpleStorage;



    function CreateSimple() public {
        simpleStorage.push(new SimpleStorage()); // 部署一个新的合约
    }

    //  合约之间可以相互组合调用，和其他的语言一样，组合，聚合，依赖，等关系
    function SFStore(uint _sStoreIndex,uint _SFNum) public{
        simpleStorage[_sStoreIndex].store(_SFNum);

    }
}
//  合约也和其他面向对象语言一样，继承，封装，多台

//  子合约
contract a is SimpleStorage{

    // 如果需要重写父合约的某个方法，需要添加 override 关键字
    // 父合约方法上需要添加 virtual 这个关键字
    function store(uint _num) public override{
        
    }

}
