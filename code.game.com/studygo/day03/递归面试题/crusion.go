package main

import "fmt"

//递归需要一个明确的条件退出，否则会死循环

//计算n的阶乘

func f(n int64) int64 {
	if n < 1 {
		return 1
	}
	return n * f(n-1)
}

/*
有个上台阶的面试题：
n个台阶，一次可以走1步，一次也可以走2步，共计多少种走法
*/
func taijie(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := taijie(7)
	fmt.Println(ret)

}
