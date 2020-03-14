package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) person {
	return person{
		Name: name,
		Age:  age,
	}
}

func main() {
	person := newPerson("Michal", 18)
	fmt.Printf("新人物：%#v\n", person)
}
