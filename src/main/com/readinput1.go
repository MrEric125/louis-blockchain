package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "57.12/5212/go"
	format                 = "%f / %d / %s"
)
var inputReader *bufio.Reader
var err error

func main() {

	inputFunc()

}

/*
 */
func inputFunc() {
	/**
	os.Open 读取文件
		返回的是File 的引用类型
	*/
	inputFile, inputError := os.Open("/Users/louis/workspace/louis/louis-blockchain/src/main/com/input.dat")
	if inputError != nil {
		fmt.Printf("an error ecc %s", inputError)
		return

	}
	defer inputFile.Close()
	/*
	 bufio.NewReader
	 param   io.Reader 接口数据，File 实现了io.Reader 接口,go 中实现接口的方式 不是很方便找到哪个类实现了哪个接口的哪个方法
	 return  bufio.Reader struct
	*/
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		fmt.Printf("the input was :%s \n", inputString)
		if readError == io.EOF {
			return
		}

	}
}

func input2() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
		fmt.Print("input end")
	}
	return

}
