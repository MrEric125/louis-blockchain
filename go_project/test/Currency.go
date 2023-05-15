package main

import (
	"fmt"
	"sort"
)

// Currency 是一个自定义的 int 类型
type Currency int

const (
	USD Currency = iota // 美元

	EUR // 欧元

	GBP // 英镑

	RMB // 人民币
)

/**
* 内置的 一些函数： make len cap new delete append
 */
func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	months := [...]string{1: "January", 12: "December"}

	Q2 := months[4:7]

	sumer := months[5:8]

	fmt.Println(Q2, len(Q2), cap(Q2))

	ages := make(map[string]int)

	ages["alice"] = 31
	ages["charlie"] = 34

	fmt.Println(ages["alice"])

	delete(ages, "alice")

	var names []string

	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

}
