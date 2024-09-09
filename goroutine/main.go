package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("HELLO ", i)
}

func sendData(ch chan int) {
	fmt.Println("Sending data...")
	time.Sleep(5 * time.Second)
	ch <- 42 // 发送操作将会阻塞，直到有人接收。
	fmt.Println("Data sent!")
}

func receiveData(ch chan int) {
	fmt.Println("Receiving data...")
	<-ch // 接收操作将会阻塞，直到有数据可以接收。
	fmt.Println("Data received!")
}

func f2(c chan int) {
	for {
		v, ok := <-c
		if !ok {
			fmt.Println("chan status is end!")
			break
		}
		fmt.Println("receive value is ", v, " chan status is ", ok)
	}
}

func Produces() <-chan int {
	c := make(chan int, 2)
	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func Consumer(c <-chan int) {
	for v := range c {
		fmt.Println("consumer value ", v)
	}
}

func main() {
	//ch := make(chan int) // 创建一个未缓冲的通道
	//go sendData(ch)
	//go receiveData(ch)
	//time.Sleep(10 * time.Second) // 让main goroutine等待一段时间，以确保其他goroutines完成执行。

	//c := make(chan int, 6)
	//c <- 10
	//c <- 20
	//c <- 30
	//c <- 40
	//c <- 50
	//c <- 60
	//go f2(c)
	//time.Sleep(time.Second * 60)
	//close(c)

	//c := Produces()
	//Consumer(c)
	//time.Sleep(time.Second * 10)

	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
