package main

import "fmt"

//需要一个实现:
// 1、它保存了一些数据---->结构体的字段
// 2、它有三个功能----->结构体的方法

//创建管理者
type studentMgr struct {
	allStudent map[int]*student
}

//构造函数
func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func (s studentMgr) showAllStudent() {
	for k, v := range s.allStudent {
		fmt.Printf("学生学号是：%d,姓名是：%s\n", k, v.name)

	}
	fmt.Println("----------分割线-----------")
}

func (s studentMgr) addStudent() {
	var (
		id   int
		name string
	)
	fmt.Println("请输入学号：")
	fmt.Scanln(&id)

	if _, ok := s.allStudent[id]; !ok {
		fmt.Println("请输入姓名：")
		fmt.Scanln(&name)
		newStu := newStudent(id, name)
		s.allStudent[id] = newStu
		fmt.Println("添加成功！！")
	} else {
		fmt.Println("该学生的学号已存在，请重新输入!")
		return
	}

}

func (s studentMgr) deleteStudent() {
	var id int
	fmt.Println("请输入删除的学生学号：")
	fmt.Scanln(&id)
	if _, ok := s.allStudent[id]; ok {
		delete(s.allStudent, id)
		fmt.Println("删除成功！！")
	} else {
		fmt.Println("该学生的学号不存在，请重新输入!")
		fmt.Println("----------分割线-----------")
		return
	}

}

func (s studentMgr) updateStudent() {
	//1、获取用户输入
	//2、展示该学号对应的学生信息，如果没有提示查无此人
	//3、请输入修改后的学生名
	//4、更新学生的姓名
	var id int
	fmt.Println("请输入学号：")
	fmt.Scanln(&id)
	stu, ok := s.allStudent[id]
	if !ok {
		fmt.Println("查无此人!")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号--->%d，姓名--->%s\n", stu.id, stu.name)
	fmt.Println("请输入学生的姓名：")
	var newName string
	fmt.Scanln(&newName)
	stu.name = newName
	s.allStudent[id] = stu
	fmt.Println("更新成功！！")
}
