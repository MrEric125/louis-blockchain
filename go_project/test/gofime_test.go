package main

import (
	"fmt"
	"os"
	"testing"
)

var (
	newFile *os.File
	err error
)

func TestNewFile(t *testing.T) {

	exist := IsFileExist("test.txt")

	fmt.Println("文件是否存在", exist)


}

func IsFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}