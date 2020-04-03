package main

import (
	"fmt"
	"strconv"
	"sync"
)

//go语言中的内置map不是并发安全的

//var m=make(map[string]int)

var m2 = sync.Map{} //开箱即用，无需初始化

//func getKey(key string) int {
//	return m[key]
//}
//
//func setKey(key string,value int)  {
//	m[key]=value
//}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 6800; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			//setKey(key,n)
			m2.Store(key, n)         //使用store方法设置键值对
			value, _ := m2.Load(key) //使用load方法根据key取值
			fmt.Printf("k=%s,v=%d\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
