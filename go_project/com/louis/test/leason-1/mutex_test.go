package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {

	lock := &sync.RWMutex{}

	lock.Lock()

	for i := 0; i < 10; i++ {
		go func(i int) {

			fmt.Printf("第 %d 个携程准备执行\n", i)

			lock.RLock()
			fmt.Printf("第 %d 个协程获得读锁, sleep 1s 后，释放锁\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()

		}(i)
	}
	time.Sleep(time.Second * 2)

	fmt.Println("准备释放写锁，读锁不再阻塞")

	lock.Unlock()

	lock.Lock()

	fmt.Println("程序退出")

	lock.Unlock()

}

func TestChannel(t *testing.T) {
	pipeline := make(chan int)

	go func() {
		fmt.Println("准备发送数据: 100")
		pipeline <- 100

	}()

	go func() {
		num := <-pipeline
		fmt.Printf("接受的消息:%d \n", num)

	}()

	// 只读信道
	type Receiver = <-chan int
	// var receive Receiver = pipeline

	// 只写信道
	type Sender = chan<- int // 关键代码：定义别名类型
	// var sender Sender = pipeline

	go func() {
		var sender Sender = pipeline
		fmt.Println("准备发送数据: 200")
		sender <- 200
	}()

	go func() {
		var receiver Receiver = pipeline
		num := <-receiver
		fmt.Printf("接收到的数据是: %d\n", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)

}

// 简单的流水线模式
func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out

}
func TestPipeline(t *testing.T) {
	in := producer(1, 2, 3, 4)
	ch := square(in)

	// consumer
	for ret := range ch {
		fmt.Printf("%3d", ret)
	}
	fmt.Println()

}
