package main

import (
	"fmt"
	"strings"
)

func main() {
	//本来\就具有特殊含义，
	path := "\"C:\\workspace\\src\\code.game.com\\studygo\\day01\""
	fmt.Println(path)
	s := "I'm ok,你呢？영웅연맹"
	fmt.Println(s)

	//多行字符串,使用反引号``，esc键下面的符号
	s2 := `
	鹅鹅鹅，曲项向天歌
	白毛浮绿水，红掌拨清波
	`
	fmt.Println("s2=", s2)

	path2 := `C:\workspace\src\code.game.com\studygo\day01`
	fmt.Println("path2=", path2)

	//分割
	ret := strings.Split(path2, "\\")
	fmt.Println("ret=", ret)
	//拼接
	s3 := strings.Join(ret, "+")
	fmt.Println("s3= ", s3)
	//len()求得是byte字节的数量
	//byte和rune类型
	/*go语言中为了处理非ASCII码类型的字符，定义了新的rune类型
	字符分为两种：
	1、unit8类型，或者叫做byte型，代表了ASCII码的一个字符
	2、rune类型，代表一个UTF-8字符，包含中文、日文、韩文等其他语言;rune类型本质是int32
	*/
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
	//字符串修改
	ss := "白萝卜"
	sss := []rune(ss) //把字符串强制转换成一个rune切片
	sss[0] = '红'
	fmt.Println(string(sss))

}
