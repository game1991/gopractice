package main

import "fmt"

//注册姓名
func register(name string) {
	fmt.Println("Hello", name)
}

//函数作为参数
func two(f func(string), name string) {
	f(name)
}

//函数作为返回值
func three() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}
//无参无返回值
func low(f func()){
	f()
}
//闭包
func bi(f func(string),name string)func(){
	return func(){
		f(name)
	}
}

func main() {
	two(register, "小马王")
	ret := three()
	fmt.Printf("%T\n", ret)
	sum := ret(10, 20)
	fmt.Println("sum=", sum)

	//闭包示例
	fc:=bi(register,"闭包的内容")
	low(fc)
}
