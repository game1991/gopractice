package main

import (
	"fmt"
	"sync"
	"time"
)

//读远远大于写的场景，适用于读写锁；例如读取数据库，读写锁常用

var (
	x1  int
	wg1 sync.WaitGroup
	//lock1 sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	defer wg1.Done()
	//lock1.Lock()
	rwLock.RLock()
	fmt.Println(x1)
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	//lock1.Unlock()
}

func write() {
	defer wg1.Done()
	//lock1.Lock()
	rwLock.Lock()
	x1 += 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
	//lock1.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg1.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg1.Add(1)
	}
	wg1.Wait()
	fmt.Println(time.Now().Sub(start)) //相隔多少时间
}
