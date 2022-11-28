package main

import (
	"fmt"
	"testing"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func TestValue(t *testing.T) {
	var b=[4]int{12,23,43}

	b =[4]int{1:10,2:30,0:20}
	for  i,i2 := range b {
		fmt.Printf("变量内容 %d %d  %x \n",i,i2,&i2)

	}

	var book=Books{"title","author","subject",123}

	fmt.Println(book)
	book.author="zhangsan"
	fmt.Println(book)

	var slice1 = make([]int , 1, 5)

	fmt.Println(slice1)




}
