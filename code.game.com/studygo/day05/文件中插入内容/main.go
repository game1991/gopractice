package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func CreateFile() string{
	//从终端读取内容
	fmt.Println("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s", s)

	file, err := os.OpenFile("./studygo/day05/sb.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("create file failed:", err.Error())
		return ""
	}
	defer file.Close()

	file.WriteString(s)

	return s

}

func InsertFile() {
	//打开要操作的文件
	file, err := os.OpenFile("./studygo/day05/sb.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("open file failed:", err.Error())
		return
	}

	//因为没有办法直接在文件中插入内容，需要借助一个临时文件
	tmpfile, err := os.OpenFile("./studygo/day05/sb.tmp",
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("create tmpfile failed:", err.Error())
		return
	}
	//读取源文件写入临时文件
	var ret [1]byte
	n, err := file.Read(ret[:])
	if err != nil {
		log.Fatal("read from file failed:", err.Error())
		return
	}
	//写入临时文件
	tmpfile.Write(ret[:n])

	//再写入要插入的内容
	//file.Seek(2, 0) //光标移动到b
	s :=[]byte{'c'}
	tmpfile.Write(s)
	//把源文件后续的内容也写入临时文件
	var x [1024]byte
	for {
		i, err := file.Read(x[:])
		if err==io.EOF{
			tmpfile.Write(x[:i])
			break
		}
		if err!=nil{
			log.Fatal("read file another part failed!",err.Error())
			return
		}
		tmpfile.Write(x[:i])
	}
	//重命名tmpfile替换源文件
	file.Close()
	tmpfile.Close()
	os.Rename("./studygo/day05/sb.tmp","./studygo/day05/sb.txt")

}

func main() {
	InsertFile()
}
