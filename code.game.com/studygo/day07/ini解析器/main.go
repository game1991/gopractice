package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadini(fileName string, data interface{}) (err error) {
	/*参数校验
	1、传进来的data必须是指针类型，因为需要在函数中对其赋值；
	2、传进来的data类型必须是结构体指针类型，因为配置文件中的各种键值对都需要赋值给结构体的字段)
	*/
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data should be a Pointer") //格式化输出后返回一个error类型
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be struct pointer")
		return
	}
	//1.读文件得到字节类型数据
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(bytes), "\n")
	fmt.Printf("%#v\n", lineSlice)
	//2.一行一行读数据
	var structName string
	for index, line := range lineSlice {
		//去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		//遇到空行就跳过
		if len(line) == 0 {
			continue
		}
		//2.1如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//2.2如果是[开头的就表示节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d,syntax error", index+1)
				return
			}
			//把这一行的首尾[]去掉，取到中间的内容并且去除空格
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			fmt.Printf("sectionName=%#v\n", sectionName)
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d,syntax error", index+1)
				return
			}
			//根据sectionName去data里面利用反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			//2.3如果不是[开头就是=键值对

			//2.3.1以等号分割这一行，等号左边是key，右边是value
			//首先筛选出不符合格式的line去除
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", index)
				return
			}
			idx := strings.Index(line, "=")
			key := strings.TrimSpace(line[:idx])
			value := strings.TrimSpace(line[idx+1:])
			//2.3.2根据structName去data里面把对应的嵌套结构体给取出来,因为是对其赋值，所以使用Valueof
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是结构体", structName)
				return
			}
			var fieldName string
			//2.3.3遍历结构体的每一个字段，判断tag是否等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) //tag信息是存储在类型信息中的
				if field.Tag.Get("ini") == key {
					//找到了结构体对应的字段
					fieldName = field.Name
				}
			}
			//2.3.4如果key=tag，就给这个字段赋值
			//根据fileName去取出这个字段对其赋值
			if len(fieldName) == 0 {
				//在结构体中找不到对应的字段直接跳过
				continue
			}
			fieldObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fieldObj.Type().Kind())
			switch fieldObj.Type().Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", index+1)
					return
				}
				fieldObj.SetFloat(valueFloat)
			}

		}

	}
	return
}

func main() {
	var cfg Config
	//传入的参数为指针是因为需要修改结构体成员的赋值
	err := loadini("./studygo/day07/ini解析器/config.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini file failed,err:%v\n", err)
		return
	}
	fmt.Printf("cfg=%#v\n", cfg)
	//fmt.Println(cfg.MysqlConfig.Address,cfg.MysqlConfig.Port,cfg.MysqlConfig.Username,cfg.MysqlConfig.Password)
}
