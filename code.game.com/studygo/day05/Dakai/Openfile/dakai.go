package openfile

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Dakai ...
func Dakai() (file *os.File, err error) {
	file, err = os.OpenFile("./xx.txt", os.O_WRONLY|os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v\n", err)
		return
	}
	return file, nil
}

//WriteDemo ...
func WriteDemo() {
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
