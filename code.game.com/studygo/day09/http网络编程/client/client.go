package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)
//公用一个client适用于请求比较频繁
var (
	client =http.Client{
		Transport:     &http.Transport{DisableKeepAlives:false},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5,
	}
)

func main() {
	/*resp, err := http.Get("127.0.0.1:9090/xxx?name=甘磊&age=28")
	if err != nil {
		log.Fatal("get url failed,err:", err)
		return
	}*/
	urlObj, _ := url.Parse("127.0.0.1:9090/xxx")
	data := url.Values{} //url Values
	data.Set("name", "甘磊")
	data.Set("age", "28")
	urlEncode := data.Encode() //url encode之后的URL
	fmt.Println(urlEncode)
	urlObj.RawQuery = urlEncode
	request, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		log.Fatalf("http Newrequest failed,err:%#v\n", err)
		return
	}
	/*  使用默认的Client
	resp, err1 := http.DefaultClient.Do(request)
	if err1 != nil {
		log.Fatalf("httpDefaultClient Do request failed,err:%#v\n", err1)
		return
	}
	*/

	//禁用KeepAlive的client 实现一个短连接,适用于请求不是特别频繁的情景,使用完就关闭释放
	/*
	tr:=&http.Transport{DisableKeepAlives:true}
		client:=http.Client{Transport:tr}
	*/

	resp, err1 := client.Do(request)
	if err1 != nil {
		log.Fatalf("httpDefaultClient Do request failed,err:%#v\n", err1)
		return
	}
	defer resp.Body.Close() //程序使用完response后必须关闭回复的主体

	//从resp中把服务端返回的数据读出来
	/*
		var data []byte
		resp.Body.Read(data)
		resp.Body.Close()
	*/
	if all, err := ioutil.ReadAll(resp.Body); err != nil { //在客户端读出服务端返回响应的Body
		log.Fatal("read resp.Body failed,err:", err)
		return
	} else {
		fmt.Println(string(all))
	}

}
