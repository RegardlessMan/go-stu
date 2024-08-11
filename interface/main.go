package main

import (
	"fmt"
	"stu/Cdemo"
)

type car interface {
	stop()
	start(add int)
}

type bc struct {
	brand    string
	capacity int
}

type bm struct {
	brand    string
	capacity int
}

func (b bc) stop() {
	fmt.Println(b.brand, "停止")
}
func (d bm) stop() {
	fmt.Println(d.brand, "停止")
}
func (b bc) start(add int) {
	fmt.Println(b.brand, "启动")
	b.capacity += add
}
func (d bm) start(add int) {
	fmt.Println(d.brand, "启动")
	d.capacity += add
}

func start(c car, add int) {
	c.start(add)
}

// 类型断言
func test(a any) {
	fmt.Println("----开始类型断言测试------------")

	switch t := a.(type) {
	case string:
		fmt.Println("this is string", t)
	case int:
		fmt.Println("this is int", t)
	}

	fmt.Println("----结束类型断言测试------------")
}

// 定义一个函数，该函数接受一个空接口类型的参数
func processValue(v any) {
	switch value := v.(type) {
	case int:
		fmt.Printf("Received an integer: %d\n", value)
	case string:
		fmt.Printf("Received a string: %s\n", value)
	case float64:
		fmt.Printf("Received a float64: %f\n", value)
	default:
		fmt.Printf("Received a value of unknown type: %T\n", value)
	}
}

// 接口学习相关
func main() {
	//bc := bc{
	//	brand:    "奔驰",
	//	capacity: 200,
	//}
	//bm := &bm{
	//	brand:    "宝马",
	//	capacity: 300,
	//}
	//start(bc, 1000)
	//fmt.Println(bc)
	//start(bm, 1000)
	//fmt.Println(bm)
	test(100)
	fmt.Println(Cdemo.Pi)
}
