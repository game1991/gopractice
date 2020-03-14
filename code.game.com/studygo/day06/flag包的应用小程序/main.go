package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	NumStr  = "0123456789"
	charStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

/*
   实现一个密码生成工具，支持以下功能：
   1、用户可以通过-l 指定生成密码的长度
   2、用户可以通过-t 指定生成密码的字符集，例如：
   -t num 生成全是数字的密码
   -t char 生成的全是英文字符的密码
   -t mix 生成包含数字和英文字符的密码
   -t advance 生成包含数字、英文以及特殊字符的密码
*/

var (
	length  int
	charset string
)

//解析的方法
func parseArgs() {
	flag.IntVar(&length, "l", 10, "-l是生成密码的长度参数")
	flag.StringVar(&charset, "t", "num", "-t是指定字符集 num是数字...")
	flag.Parse()
}

//生成密码的方法
func myPassword() string {
	//切片存
	password := make([]byte, length, length)
	var res string
	switch charset {
	case "num":
		res = NumStr
	case "char":
		res = charStr
	case "mix":
		res = fmt.Sprintf("%s%s", NumStr, charStr)
	case "advance":
		res = fmt.Sprintf("%s%s%s", NumStr, charStr, SpecStr)
	default:
		res = NumStr
	}

	//根据长度赋值:
	for i := 0; i < length; i++ {
		index := rand.Intn(len(res))
		password[i] = res[index]
	}
	return string(password)
}

func main() {

	rand.Seed(time.Now().UnixNano())
	parseArgs()
	password := myPassword()
	fmt.Println(password)
}
