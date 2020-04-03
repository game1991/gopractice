package main

import (
	"encoding/json"
	"fmt"
)

//json
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"周林","age":9000}`
	var p Person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}
