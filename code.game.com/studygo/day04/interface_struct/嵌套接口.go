package main

import "fmt"

// WashingMachine ...洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var x WashingMachine

	h := haier{}

	x=h

	x.dry()
	x.wash()
}

/*x.(T)该语法返回两个参数，
第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，
若为true则表示断言成功，为false则表示断言失败。
*/
