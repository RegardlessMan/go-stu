package main

import (
	"fmt"
	"reflect"
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

// 反射测试函数
func reflectType(r any) {
	t := reflect.TypeOf(r)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(r any) {
	v := reflect.ValueOf(r)
	fmt.Printf("Value is %v\n", v)
	t := v.Kind()
	if reflect.Int == t {
		fmt.Printf("good param is Int\n")
	}
}

func reflectSetValue(value any) {
	v := reflect.ValueOf(value)
	v.Elem().SetInt(200)
}

// 接口学习相关
func main() {
	//type myInt int
	//var a *float32 // 指针
	//var b myInt    // 自定义类型
	//var c rune     // 类型别名
	//reflectType(a) // type: kind:ptr
	//reflectType(b) // type:myInt kind:int64
	//reflectType(c) // type:int32 kind:int32
	//
	//type person struct {
	//	name string
	//	age  int
	//}
	//type book struct{ title string }
	//var d = person{
	//	name: "沙河小王子",
	//	age:  18,
	//}
	//var e = book{title: "《跟小王子学Go语言》"}
	//reflectType(d) // type:person kind:struct
	//reflectType(e) // type:book kind:struct
	a := 1
	reflectValue(a)
	reflectSetValue(&a)
	fmt.Println(a)
}
