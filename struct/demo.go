package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func newPerson(name string, age int, sex int) *Person {
	return &Person{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
}

func (p *Person) setName(name string) {
	p.Name = name
}

type Human struct {
	Cur  Person
	Desc string
}

func main() {
	h1 := &Human{
		Cur:  Person{},
		Desc: "null",
	}
	h1.Cur.Name = "test"
	h1.Cur.Age = 1
	h1.Cur.Sex = 0
	fmt.Println(h1)
	data, err := json.Marshal(h1)
	if err != nil {
		fmt.Println("json marshal fail!")
		return
	}
	fmt.Printf("json %s\n", data)
	h2 := &Human{}
	err = json.Unmarshal([]byte(data), h2)
	if err != nil {
		return
	}
	fmt.Printf("%#v", h2)
}
