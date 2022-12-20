package main

import (
	"fmt"
	"testing"
)

// 创建测试方法如果需要生效，需要将这个类名结尾为 (_test.go)  创建的方法必须是(Test)开头的
func abcd() {

}

type myint int

type Book struct {
	name string
	auth string
}

type treeNode struct {
	value       int
	left, right *treeNode
}

// 这种情况，传递的是值对象，如果在这个方法中修改了book对象，外部对象是不会受到影响的
func changeBook(book Book) {
	book.name = "java"
}

// 传递的是引用类型，如果方法内修改了这个对象，那么外层对象也会受到影响
func changeBook2(book *Book) {

	book.name = "go"
}

func TestV2(t *testing.T) {
	var a myint = 10
	fmt.Println("a=", a)             //a= 10
	fmt.Printf("type of a =%T\n", a) //type of a =main.myint

	var book1 Book
	book1.name = "Golang"
	book1.auth = "zhangsan"
	fmt.Printf("%v\n", book1) //{Golang zhangsan}

	changeBook(book1)
	fmt.Println(book1.name) //Golang
	changeBook2(&book1)
	fmt.Println(book1.name) //Java
}

func TestTreeNode(t *testing.T) {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	nodes := []treeNode{
		{3, nil, nil},
		{},
		{6, nil, &root},
	} //这里[]treeNode前面不用加“&”，因为切片是引用类型
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000004078}]

}