package main

import (
	"fmt"
	"unicode"
)

func main() {
	//判断字符串中汉字的数量
	/*1、依次拿到字符串中的字符
	  2、判断当前这个字符是不是汉字
	  3、把汉字出现的次数累加得到最终结果
	*/
	s := "hello小马王마왕"
	var count int
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++  
		}
	}
	fmt.Println(count)
}
