package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(file)
}

func f2(w http.ResponseWriter, r *http.Request) {
	//对于GET请求,参数都放在URL上(query param),请求体里面是没有数据的
	fmt.Println(r.URL.String())
	queryParam := r.URL.Query() //自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)               //获取方法,例如get/post
	fmt.Println(ioutil.ReadAll(r.Body)) //在服务端打印客户端发来请求的Body
	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/hello", f1)
	http.HandleFunc("//xxx", f2)

	http.ListenAndServe("0.0.0.0:9090", nil) //0.0.0.0代表全网都可以访问
}
