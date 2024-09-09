package main

import (
	"fmt"
	"time"
)

func timeSub(t time.Time) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed!")
		return
	}
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", "2024-09-09 14:43:32", loc)
	if err != nil {
		fmt.Println("parse is failed!")
	}
	fmt.Println(t1)
	sub := t1.Sub(t)
	fmt.Println(sub)

}

func main() {
	now := time.Now()
	timeSub(now)
}
