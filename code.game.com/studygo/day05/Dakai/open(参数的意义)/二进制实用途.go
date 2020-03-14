package main

import "fmt"

const (
	eat   int = 4
	sleep int = 2
	da    int = 1
)

//111
/*
左边的代表吃饭 100
中间的代表睡觉 010
右边的代表打豆豆 001
*/

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | da)         //101
	f(eat | sleep | da) //111
}
