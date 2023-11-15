package main

import (
	"fmt"
	"testing"
)

func TestForeach(t *testing.T) {
	var counts int = 10
	for i := 0; i < counts; i++ {
		fmt.Print(i)
	}
	v := make(map[string]int)
	for _, vab := range v {
		fmt.Print(vab)
	}
	fmt.Println("")
	fmt.Println(fib(6))
	medals := []string{"1", "2", "3", "4", "5"}
	fmt.Println(medals)
	fmt.Println(medals[1:3])

	// 对应地址，以及对应值的获取
	var age int = 10
	fmt.Println("age 对应地址为：", &age)
	fmt.Println("age 对应值为：", age)
	var add *int = &age
	fmt.Println("age 对应地址为：", &add)
	fmt.Println("age 对应值为：", *add)


}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y

	}
	return x
}
