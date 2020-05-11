package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	getFunc()
	postFunc()
}


func postFunc(){
	resp, err := http.PostForm("http://127.0.0.1:10666/user", url.Values{
		"name":      []string{"zhangsan"},
		"telephone": []string{"18888888888"},
		"email":     []string{"99@mail.com"},
		"password":  []string{"123456"},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)
	fmt.Println(err, string(bts))
}

func getFunc(){
    resp, err := http.Get("http://127.0.0.1:10666/user/zhangsan")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)
	fmt.Println(err, string(bts))
}