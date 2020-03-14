package main

import (
	"flag"
	"fmt"
)

func main() {
	//flag包操作的事情等同于下面的代码
	/*for index,value:=range os.Args{
		fmt.Printf("%d--->%v\n",index,value)
	}*/
	/*
	   参数1：传的参数key
	   参数2：默认值
	   参数3：说明
	*/
	//如果命令行没有输入指令，显示默认值
	/*i := flag.String("name", "无名氏", "名字")
	i2 := flag.Int("age", 0, "年龄")
	i3 := flag.Float64("money", 0, "财富")
	i4 := flag.Bool("isstupid", true, "是否愚蠢")
	//解析参数
	flag.Parse()
	fmt.Println(*i, *i2, *i3, *i4)*/

	fmt.Println("-------------分割线-----------------")

	var (
		name string
		age int
		money float64
		isstupid bool
	)
	flag.StringVar(&name,"name", "无名氏", "名字")
	flag.IntVar(&age,"age", 0, "年龄")
	flag.Float64Var(&money,"money", 0, "财富")
	flag.BoolVar(&isstupid,"isstupid", true, "是否愚蠢")

	flag.Parse()
	fmt.Println(name,age,money,isstupid)


}
