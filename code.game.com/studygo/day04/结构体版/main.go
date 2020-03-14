package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

func showMenu() {
	fmt.Println("-------------欢迎来到学生管理系统！-------------")
	fmt.Println(`----------------
		#1.查看所有学生#
		#2.新增学生    #
		#3.删除学生    #
		#4.修改学生    #
		#5.退出        #
		----------------`)
}

func main() {
	smr := studentMgr{
		allStudent: make(map[int]*student),
	}
	for {

		fmt.Println("请输入数字实现对应的操作(按0查看帮助):")
		var InputNum int
		fmt.Scanln(&InputNum)
		fmt.Printf("你选择的是第%d的选项\n", InputNum)
		switch InputNum {
		case 0:
			showMenu()
		case 1:
			smr.showAllStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.deleteStudent()
		case 4:
			smr.updateStudent()
		case 5:
			os.Exit(1)
			fmt.Println("----------退出了！-----------")
		default:
			fmt.Println("滚！！！")
		}
	}
}
