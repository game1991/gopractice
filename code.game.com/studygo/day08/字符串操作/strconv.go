package main

import (
	"fmt"
	"strconv"
)

func main() {
	//从字符串中解析出对应的数据
	str := "10000"
	ret1, _ := strconv.ParseInt(str, 10, 64) //解析字符串转换成10进制的int64
	fmt.Printf("%#v %T\n", ret1, ret1)

	//把数字转换成字符串
	i := int32(97)
	//ret2:=string(i)  //"a"  对应的是其ASCII码
	ret2 := fmt.Sprintf("%d", i) //"97"
	fmt.Printf("%#v\n", ret2)
	//从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

}
