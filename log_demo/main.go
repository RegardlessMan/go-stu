package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	MidGet()
}

func MidGet() {
	apiUrl := "http://127.0.0.1:8080/ping"
	//url param
	data := url.Values{}
	data.Set("name", "小王子")
	data.Set("age", "18")

	u, _ := url.ParseRequestURI(apiUrl)

	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))

}

func LogStu() {
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

func BasicGet() {
	resp, err := http.Get("http://127.0.0.1:8080/ping")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))
}
