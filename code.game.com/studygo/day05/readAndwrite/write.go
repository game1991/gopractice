package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//WriteDemo ...
func WriteDemo() {
	var s string
	fmt.Println("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是：%s\n", s)

	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(s)

}

//WriteDemo2 ...
func WriteDemo2() {
	//从终端读取内容
	fmt.Println("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
	//打开文件，如果没有就新建文件
	file, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	defer file.Close()

	wr := bufio.NewWriter(file)
	wr.WriteString(s) //将数据写入缓存中
	wr.Flush()        //将缓存中的内容写入文件
}

//WriteDemo3 ...
func WriteDemo3() {
	//从终端读取内容
	fmt.Println("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
	//写入文件
	err := ioutil.WriteFile("./xx.txt", []byte(s), 0777)
	if err != nil {
		fmt.Printf("write file failed,err is: %v\n", err)
		return
	}
}
