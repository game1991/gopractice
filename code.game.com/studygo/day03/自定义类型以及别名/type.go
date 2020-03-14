package main

import "fmt"

//自定义类型
type myInt int

//YourInt is "别名"
type YourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("n的类型是：%T\n", n)
	var m YourInt
	m = 200
	fmt.Println(m)
	fmt.Printf("m的类型是：%T\n", m)

	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("c的值是%c\nc的类型是：%T\n", c, c)

}
