package main

import (
	"fmt"
	"strings"
)

func main() {
	//求出字符串中单词出现的次数
	s := "how do you do"
	//1、处理字符串得到单词
	s1 := strings.Split(s, " ")
	//2、遍历切片存储到map
	m := make(map[string]int)
	for _, w := range s1 {
		//1、如果原来的map中不存在w这个key,那么出现次数等于1
		if _, ok := m[w]; !ok {
			m[w] = 1
		} else { //2、如果map中存在这个key，那么出现次数加1
			m[w]++
		}
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

}
