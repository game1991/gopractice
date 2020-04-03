package main

import (
	"fmt"
	"time"
)

//goroutine 池

/*
控制goroutine的数量，防止goroutine泄露或者内存暴涨

*/
//work pool是用来控制性能，通过goroutine的适量增加调整，达到最快性能，一般按照64,128或者更大去增加goroutine

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
		notifyCh <- struct{}{} //匿名函数结构体实例化，这样的方式更节省空间，一般用来通知用
	}

}

/*
type cat struct{} //声明类型
var c1=cat{} //实例化
*/

var (
	notifyCh = make(chan struct{}, 5)
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启5个任务
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) //开启这三个goroutine就是协程池
	}

	go func() {
		for i := 0; i < 5; i++ {
			<-notifyCh
		}
		close(results)
	}()
	// 输出结果
	//for a := 1; a <= 5; a++ {
	//	<-results
	//}

	//for{
	//	x,ok:=<-results
	//	if !ok{
	//		break  //什么时候ok=false? results通道被关闭的时候
	//	}
	//	fmt.Println(x)
	//}

	//输出结果
	go func() {
		for x := range results { //for range模式更好一点 当然需要注意死锁情况，需要关闭通道
			fmt.Println(x)
		}
	}()

}
