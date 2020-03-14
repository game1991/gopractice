package main

import "fmt"

func main() {
	//判断一个map的key是否存在
	m := make(map[string]interface{})
	m["foo"] = 007
	m["zoo"] = 001
	m["hello"] = "你好"
	m["Are you ok?"]="你好吗？"

	v, ok := m["Are you ok?"]
	if !ok {
		fmt.Println("map the key not exist!")
		return
	}
	fmt.Printf("map the key is \"Are you ok?\",the value is %#v\n",v)
}
