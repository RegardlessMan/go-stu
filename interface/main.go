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
func testType(a any) {
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

func test() {
	v := 0
	r := reflect.ValueOf(v)
	r.IsValid()
	print(r.Kind() == reflect.Int)
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func testReflect1() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func (s student) Study() string {
	fmt.Println("Study")
	return "好好学习，天天向上"
}

func (s student) Sleep() string {
	fmt.Println("Sleep")
	return "我睡觉啦"
}

func testReflect2(x any) {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		methodType := method.Type()
		fmt.Printf("method name:%s index:%d type:%s\n", t.Method(i).Name, i, methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		method.Call(args)
	}

}

// 接口学习相关
func main() {
	//testReflect1()
	testReflect2(student{})
}
