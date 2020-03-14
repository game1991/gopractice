package main

import "fmt"

//接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
//在go语言中interface是一种类型，一种抽象的类型

/*interface是一组method的集合，是duck-type programming的一种体现。
接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。
不关心属性（数据），只关心行为（方法）。
*/

//animal
type animal interface {
	move()
	eat(string)
}

//Sayer ...
type Sayer interface {
	say() //只要实现了say()方法的变量都是speaker类型，方法签名
}

//Cat ...
type Cat struct {
	name string
	feet int
}

// //使用值接收者实现了接口的所有方法
// func (c Cat) move() {
// 	fmt.Println("走猫步")
// }
// func (c Cat) eat(food string) {
// 	fmt.Printf("猫吃%s\n", food)
// }

//使用指针接收者实现了接口的所有方法
func (c *Cat) move() {
	fmt.Println("走猫步")
}
func (c *Cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func (c Cat) say() {
	fmt.Println("喵喵喵！")
}

//Dog ...
type Dog struct {
}

func (d Dog) say() {
	fmt.Println("汪汪汪！")
}

func da(x Sayer) {
	//传进来一个参数，传进来什么就打什么
	x.say() //挨打了就会叫
}

func main() {

	var (
		c1 Cat
		// d1 Dog
		// x  Sayer //定义一个接口类型的变量x
		a animal
	)
	c1 = Cat{"淘气", 3}
	c2 := &Cat{"萌萌", 5}
	// da(x)
	// x = c1
	// x = d1
	// fmt.Println(x)
	// da(c1)
	// da(d1)
	a = &c1 //实现animal这个接口类型的是指针类型
	fmt.Println(a)
	a = c2
	fmt.Println(a)

	/*
		1、使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
		2、使用指针接收者实现接口只能存结构体指针类型的变量
	*/
}
