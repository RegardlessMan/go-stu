package main

import "fmt"

func main() {
	//m := make(map[int]string, 100)
	//for i := 0; i < 100; i++ {
	//	m[i] = fmt.Sprintf("stu%02d", i)
	//}
	//nums := make([]int, 0, 100)
	//for k := range m {
	//	nums = append(nums, k)
	//}
	//sort.Ints(nums)
	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(m[nums[i]])
	//}
	sliceMap()
}

func sliceMap() {
	//var mapSlice = make([]map[string]string, 3)
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	//fmt.Println("after init")
	//// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 10)
	//mapSlice[0]["name"] = "小王子"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "沙河"
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}

	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s1 := []int{3, 4}
	//s = append(s[:1], s[2:]...)
	s = append(s[:1], s1[:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
