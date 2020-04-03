package main

import (
	"fmt"
	"sync"
)

/*

单向通道多用于函数的参数 控制这个管道只能做写入数据操作或者读取数据操作

*/

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) //关闭ch1，ch1无法写入数据，但是仍可以读取数据<-ch1
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for x := range ch1 {
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2)
	}) //确保某个操作只执行一次
	//如果不使用once，f2()函数调用两次close(ch2)，对已关闭的ch2再次操作会引起panic
}
func main() {
	//无缓冲通道也被称为同步通道
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(6)
	go f1(a) //生成数据
	for i := 0; i < 5; i++ {
		go f2(a, b) //消费数据
	}
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

}
