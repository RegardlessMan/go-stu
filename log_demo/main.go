package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file is failed")
		return
	}
	log.SetOutput(file)
	for {
		log.Println("这是一条测试日志")
		time.Sleep(time.Second * 10)
	}
}
