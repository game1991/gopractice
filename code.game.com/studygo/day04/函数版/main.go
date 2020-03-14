package main

import (
	"fmt"
	"os"
)

/*
支持查看/新增/删除学生系统
*/
type student struct {
	id   int
	name string
}

//构造函数
func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

//定义一个容器装学生,考虑到学生的名字可能重复，但是学号是唯一的，所以使用的是map[int]interface
var (
	allStudent = make(map[int]*student)
)

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学生学号是：%d,姓名是：%s\n", k, v.name)

	}
	fmt.Println("----------分割线-----------")
}

func addStudent() {
	var (
		id   int
		name string
	)
	fmt.Println("请输入学号：")
	fmt.Scanln(&id)

	if _, ok := allStudent[id]; !ok {
		fmt.Println("请输入姓名：")
		fmt.Scanln(&name)
		newStu := newStudent(id, name)
		allStudent[id] = newStu
	} else {
		fmt.Println("该学生的学号已存在，请重新输入!")
		return
	}
	fmt.Println("----------分割线-----------")
}

func deleteStudent() {
	var id int
	fmt.Println("请输入删除的学生学号：")
	fmt.Scanln(&id)
	if _, ok := allStudent[id]; ok {
		delete(allStudent, id)
	} else {
		fmt.Println("该学生的学号不存在，请重新输入!")
		fmt.Println("----------分割线-----------")
		return
	}
	fmt.Println("----------分割线-----------")
}
func main() {
	for {
		fmt.Println("欢迎来到学生管理系统！")
		fmt.Println(`----------------
#1.查看所有学生#
#2.新增学生    #
#3.删除学生    #
#4.退出        #
----------------`)
		fmt.Println("请输入数字实现对应的操作:")
		var InputNum int
		fmt.Scanln(&InputNum)
		fmt.Printf("你选择的是第%d的选项\n", InputNum)
		switch InputNum {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
			fmt.Println("----------退出了！-----------")
		default:
			fmt.Println("滚！！！")
		}
	}
}
