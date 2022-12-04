package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func TestTag(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		fmt.Print("fff")
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)

	var i interface{} = 10

	t1 := i.(int)

	fmt.Printf("t1: %v\n", t1)

	i = "hello"

	t2 := i.(string)

	fmt.Printf("t2: %v\n", t2)

	fmt.Print("=========\n")

	a := 10
	b := "hello"
	c := true
	myFunc(a, b, c)

}

func myFunc(ifaces ...interface{}) {
	for _, v := range ifaces {
		fmt.Printf("v: %v\n", v)

	}
}

type phone interface {
	call()
}

type iphone struct {
	name string
}

func (phone iphone) call() {
	fmt.Printf("call phone.name: %v\n", phone.name)
}
func (phone iphone) call_wa_chat() {
	fmt.Printf("call we_chat :phone.name: %v\n", phone.name)
}

func TestCall(t *testing.T) {
	var mobile phone
	mobile = iphone{name: "华为"}

	mobile.call()

}
