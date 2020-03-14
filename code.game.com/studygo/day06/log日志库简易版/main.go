package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
一、支持往不同的地方输出日志
二、日志分级别：
    1、Debug
    2、Trace
    3、Info
    4、Warning
    5、Error
    6、Fatal
三、日志要支持开关控制
四、完整的日志记录要包含时间、行号、文件名、日志级别、日志信息
五、日志文件要切割
    1、按文件大小切割
       1.每次记录日志之前都判断一下当前写的这个文件大小
    2、按日期切割
    3、
*/

func main() {
	file, err := os.OpenFile("./studygo/day06/log日志库简易版/日志.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	log.SetOutput(file) //设置输出到文件

	for {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 3)
	}
}
