package main

import "fmt"

//select多路复用

//走case的这几路随机执行，如果每条路都走不通，则走default

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: //case的这几路都是随机执行
			fmt.Println(x)
		case ch <- i:
		}
	}
}
