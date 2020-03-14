package main

import (
	"bufio"
	"fmt"
	"io"

	openfile "code.game.com/studygo/day05/Dakai/Openfile"
)

// bufio按行读取示例
func main() {
	// file, err := os.Open("./xx.txt")
	// if err != nil {
	// 	fmt.Println("open file failed, err:", err)
	// 	return
	// }
	// defer file.Close()

	//打开文件
	file, err := openfile.Dakai()
	if err != nil {
		fmt.Printf("打开文件出错,err:%v\n", err)
		return
	}
	defer file.Close()
	//从终端获取内容写入文件
	WriteDemo3()
	for {
		reader := bufio.NewReader(file)
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Printf("你输入的内容是：%s\n", line)
	}
}
