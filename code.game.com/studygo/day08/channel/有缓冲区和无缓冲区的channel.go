package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sandNum(ch chan<- int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(time.Second * 5) //每隔5秒往通道中发送数据
	}
}

func main() {
	ch := make(chan int, 1)
	//ch<-100     //把一个值发送到通道中
	//<-ch        //把通道中的100取出来
	go sandNum(ch)
	for {
		x, ok := <-ch //阻塞等4秒
		fmt.Println(x, ok)
		time.Sleep(time.Second) //每隔一秒取一次数据
	}
}
