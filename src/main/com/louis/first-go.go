package main

import "fmt"

//  单行注释，自己看的，程序不会执行
/**
*  多行注释，可以生成duo,一般是在全局变量或者方法头，或者类头添加
 */

func main() {

	// 定义 用var,赋值符号 =  可以用以下方式定义一组变量
	// := 自动推导，
	var (
		addr  string
		phone int
	)

	addr = "addr"
	phone = 123

	fmt.Println(addr, phone)

	var age int = 18

	fmt.Println(age)

	fmt.Printf("num:%d,内存地址：%p \n", age, &age)

	var a int = 100

	// 数据类型转换，高变量转换为低类型，会丢精度，有的不同的类型，还转换不了
	var b int = 200

	// 变量值交换
	b, a = a, b

	fmt.Println(a, b)

	// _ 下划线匿名变量，会被直接丢弃，匿名变量能被重复利用，但是局部变量在函数内不能重复，全局变量全局不能重复,局部变量和全局变量可以重复定义，使用 的时候，局部变量会覆盖全局变量
	c, _ := bfun()
	_, d := bfun()

	fmt.Println(c, d)

	const (
		iota_a = 100
		b_a    = iota
		c_a
		d_a = "haha"
	)
	print(iota_a, b_a, c_a, d_a)

	// 算数运算符
	// 关系运算符
	// 逻辑运算符
	// go run 运行流程，先编译，再运行

}

/*


**/

func bfun() (int, int) {
	return 300, 400

}
