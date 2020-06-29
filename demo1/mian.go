package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main(){

	param:=url.Values{}
	param.Set("phone","13578648970,18976543210,13077776666")
	param.Set("user","zhangsan")
	param.Set("var","variable")
	param.Set("signature",signature(param))

	fmt.Println(param.Get("signature"))
}

func signature(param url.Values) string{
	keys:=make([]string,len(param))
	for k:=range param{
		fmt.Println(k)
		keys=append(keys,k)
	}
	data := make([]string, 0, len(keys)+2)
	data = append(data, "MNBKcHo189XlWq2ebvaJB3jGTCeYVZoC")
	for i := range keys {
		data = append(data, fmt.Sprintf("%s=%s", keys[i], param.Get(keys[i])))
	}
	data = append(data, "MNBKcHo189XlWq2ebvaJB3jGTCeYVZoC")
    return strings.Join(data,"&")
}