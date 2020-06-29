//go:generate echo $GOPATH

package main

import (
	"fmt"
	"net/url"
)

//RegionArea 注册区域
type RegionArea string

//Valid 验证注册区域
func (r RegionArea) Valid() bool {
	if r == RegionAreaMainland || r == RegionAreaOther {
		return true
	}
	return false
}

//注册区域
const (
	RegionAreaMainland RegionArea = "Mainland"
	RegionAreaOther    RegionArea = "Other"
)

func main() {
	// var str RegionArea
	// str = "haha"
	// fmt.Printf("vaild is :%#v\n", str.Valid())

	v := url.Values{}
	v.Set("name", "Ava")//Set方法将key对应的值集设为只有value，它会替换掉已有的值集。
	v.Add("name","WaWa")
	v.Set("age","18")
	v.Add("age","28")
	v.Add("friend", "Zoe")
	v.Add("friend", "Sarah")
	v.Add("friend", "Jesse")

	v.Del("friend")
	
	fmt.Printf("解析后的v:%#v\n",v.Encode())
	//"name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name")) //Get会获取key对应的值集的第一个值
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"]) //获取map的值集
	fmt.Printf("v[%s]=====>%#v\n","age",v["age"])

}
