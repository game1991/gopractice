package main

import (
	"fmt"
	"sync"
)

//CSP:通过通信来共享内存，使用channel实现多个goroutine之间的通信
//例如：两个goroutine运行中，需要有数据传输，之前使用的是全局变量slice或者map去存储
// 现在是将数据放入channel存储通信

/*
练习
1、启动一个goroutine，生成100个数发生到ch1
2、启动一个goroutine从ch1中取值，计算其平方放到ch2中
3、在main中从ch2取值打印出来
*/

//使用waitGroup是为了让协程安全退出，保证程序优雅地完成协程
var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) //关闭ch1，ch1无法写入数据，但是仍可以读取数据<-ch1
}
func f2(ch1, ch2 chan int) {
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
