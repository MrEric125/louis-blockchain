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

}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y

	}
	return x
}

// 斐波那契数列
func fib2(n int) {
	R

}
