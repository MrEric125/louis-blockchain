package main

import (
	"bufio"
	"fmt"
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
	// fmt.Println("please enter you full name:")
	// fmt.Scanln(&firstName)
	// fmt.Printf("Hi %s %s!\n  ", firstName, lastName)
	// fmt.Scanf(input, format, &f, &i, &s)
	// fmt.Println("from the string we read:", f, i, s)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("the input was :%s\n\":", input)
	}

}
