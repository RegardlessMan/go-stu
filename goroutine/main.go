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

func main() {
	ch := make(chan int) // 创建一个未缓冲的通道

	go sendData(ch)
	go receiveData(ch)

	time.Sleep(10 * time.Second) // 让main goroutine等待一段时间，以确保其他goroutines完成执行。
}
