package main


/**
go 语言中实现接口的方式
 */
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
