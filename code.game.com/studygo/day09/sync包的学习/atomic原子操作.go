package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作

var (
	x2  int64
	wg2 sync.WaitGroup
	//lock2 sync.Mutex

)

func add2() {
	//lock2.Lock()
	//x2++
	//lock2.Unlock()
	atomic.AddInt64(&x2, 1)
	wg2.Done()
}

func main() {

	for i := 0; i < 100000; i++ {
		wg2.Add(1)
		go add2()
	}
	wg2.Wait()
	fmt.Println(x2)
}
