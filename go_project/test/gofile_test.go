package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

var (
	newFile *os.File
	err     error
)

func TestNewFile(t *testing.T) {

	exist := IsFileExist("test.txt")

	fmt.Println("文件是否存在", exist)

	if !exist {
		return
	}
	file, err2 := os.Open("test.txt")

	if err2 != nil {
		fmt.Println("open file failed!,err:", err2)
	}
	tmp := make([]byte, 1)

	var content []byte

	for {
		//todo  file对象是怎么知道我具体读到文件的哪个地方
		n, err2 := file.Read(tmp)

		if err2 == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err2 != nil {
			fmt.Println("read file failed ,err:", err)
			break
		}
		//fmt.Println("读取字节数",n)
		fmt.Println("读取字节数内容:", string(tmp[:n]))

		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

func TestReader(t *testing.T) {

	userFile := "test.txt"

	exist := IsFileExist(userFile)
	if !exist {
		return
	}
	file, err := os.Open(userFile)
	if err != nil {
		fmt.Println("打开文件异常：", err)
		return
	}
	defer file.Close()

	// 建立缓冲区，将文件内容放入到缓冲区
	reader := bufio.NewReader(file)

	// 循环读取文件信息
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}

func TestCreateFile(t *testing.T) {
	file := "test2.txt"

	first, err := os.Create(file)

	if err != nil {

		fmt.Println(file, err)
		return
	}
	defer first.Close()

	for i := 0; i < 10; i++ {
		first.WriteString("Just a test!\r\n")
		first.Write([]byte("Just a test!\r\n"))
	}

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
