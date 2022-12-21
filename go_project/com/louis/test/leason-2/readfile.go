package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// 读取文件

func main() {
	file, err := os.Open("/Users/louis/workspace/louis/louis-blockchain/go_project/com/louis/test/input.dat")

	if err != nil {
		print("open file error:", err.Error())
		return
	}
	// 执行完之后关闭流操作
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		// 以分隔符的方式读取
		// data, err := r.ReadString('\n')

		// 以行的方式读取
		data, _, err := r.ReadLine()

		fmt.Printf("%v", string(data))

		//  文件末尾
		if err == io.EOF {
			break
		}
		if err != nil {
			print("read file error:", err.Error())
			break

		}

	}

}

func readfile1() {
	// 读取文件内容
	file, err := os.Open("./input.dat")

	if err != nil {
		print("open file error:", err.Error())
		return
	}
	// 执行完之后关闭流操作
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		// 以分隔符的方式读取
		// data, err := r.ReadString('\n')

		// 以行的方式读取
		data, _, err := r.ReadLine()

		fmt.Printf("%v", string(data))

		//  文件末尾
		if err == io.EOF {
			break
		}
		if err != nil {
			print("read file error:", err.Error())
			break

		}

	}
}

// 使用工具
func readfile2() {

	data, err := ioutil.ReadFile("./abc.txt")
	if err != nil {
		fmt.Println("read file err:", err.Error())
		return
	}

	// 打印文件内容
	fmt.Println(string(data))
}

func Coroutine() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello coroutine %d\n", i)

			}

		}(i)
	}
	time.Sleep(time.Microsecond)
}
func CoroutineChannel() {
	var ch1 = make(chan int) // 无缓冲channel 同步
	// var ch2 = make(chan int, 2) // 有缓冲 channel 异步的

	// chan 有点类似java中的队列，但是不能直接使用，否则会报错，如果想要运行，必须要再开一个携程不停的去请求数据
	// var ch1 = make(chan int)        // 无缓冲 channel，同步
	// ch1 <- 1
	// ch1 <- 2
	//  error: all goroutines are asleep - deadlock!
	// fmt.Println(<-ch1)

	go func() {
		for {
			n := <-ch1
			fmt.Printf("n: %v\n", n)
		}

	}()
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	// 有缓冲的方式，只要缓冲区没有满就可以一直进数据，缓冲区在填满之后没有接收也会处理阻塞状态。
	// var ch2 = make(chan int,2)
	// ch2<-1
	// ch2<-2
	// fmt.Println(ch2)
	// 不加这一行的话，是可以正常运行的
	// ch2<-3           // error: all goroutines are asleep - deadlock!

}

func ReadFile()  {
	content,err:=os.ReadFile("a.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(content))

}
func openFile()  {
	file,err :=os.Open("a.txt")
	if err != nil{
		panic(err)
	}
	r:=bufio.NewReader(file)
	buf:=make([]byte,1024)
	for  {
		n,err:=r.Read(buf)
		if err !=nil && err!=io.EOF {
			panic(err)
		}
		if n== 0 {
			break
		}
		fmt.Println(string(buf[:n]))


	}


}

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Builder interface {
	Color(Color) Builder
	Build() Interface
}
type Interface interface {
	Drive() error
	Stop() error
}
