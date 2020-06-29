package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func httpPostForm() {
	var time = Format(time.Now(), "yyyyMMddHHmmss")
	var password = "xxxxx" //密码
	pwd := md5.New()
	_,err:=pwd.Write([]byte(password + time)) // 需要将密码和时间加密
	if err!=nil{
       panic(err)
	}
	var content = "【阅信】验证码888888"
	//var content = "123456";  语音验证码内容
	var name = "ceshi"                               //账号
	resp, err := http.PostForm("localhost:8080/sms", // 目标地址
		url.Values{"name": {name}, "pwd": {hex.EncodeToString(pwd.Sum(nil))},
			"content": {content}, "phone": {"13412345678"}, "subid": {GenerateNumberCode(6)}, "mttime": {time}, "rpttype": {"1"}})
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("body read failed,err is %#v\n", err)
		return // handle error
	}
	fmt.Println(byteString(body))
}
//byteString ...
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

//Format ...
func Format(t time.Time, format string) string {
	//year
	if strings.ContainsAny(format, "y") {
		year := strconv.Itoa(t.Year())
		if strings.Count(format, "yy") == 1 && strings.Count(format, "y") == 2 {
			format = strings.Replace(format, "yy", year[2:], 1)
		} else if strings.Count(format, "yyyy") == 1 && strings.Count(format, "y") == 4 {
			format = strings.Replace(format, "yyyy", year, 1)
		} else {
			panic("format year error! please 'yyyy' or 'yy'")
		}
	}

	//month
	if strings.ContainsAny(format, "M") {
		var month string
		if int(t.Month()) < 10 {
			month = "0" + strconv.Itoa(int(t.Month()))
		} else {
			month = strconv.Itoa(int(t.Month()))
		}
		if strings.Count(format, "MM") == 1 && strings.Count(format, "M") == 2 {
			format = strings.Replace(format, "MM", month, 1)
		} else {
			panic("format month error! please 'MM'")
		}
	}
	//day
	if strings.ContainsAny(format, "d") {
		var day string
		if t.Day() < 10 {
			day = "0" + strconv.Itoa(t.Day())
		} else {
			day = strconv.Itoa(t.Day())
		}
		if strings.Count(format, "dd") == 1 && strings.Count(format, "d") == 2 {
			format = strings.Replace(format, "dd", day, 1)
		} else {
			panic("format day error! please 'dd'")
		}
	}

	//hour
	if strings.ContainsAny(format, "H") {
		var hour string
		if t.Hour() < 10 {
			hour = "0" + strconv.Itoa(t.Hour())
		} else {
			hour = strconv.Itoa(t.Hour())
		}
		if strings.Count(format, "HH") == 1 && strings.Count(format, "H") == 2 {
			format = strings.Replace(format, "HH", hour, 1)
		} else {
			panic("format hour error! please 'HH'")
		}
	}
	//minute
	if strings.ContainsAny(format, "m") {
		var minute string
		if t.Minute() < 10 {
			minute = "0" + strconv.Itoa(t.Minute())
		} else {
			minute = strconv.Itoa(t.Minute())
		}
		if strings.Count(format, "mm") == 1 && strings.Count(format, "m") == 2 {
			format = strings.Replace(format, "mm", minute, 1)
		} else {
			panic("format minute error! please 'mm'")
		}
	}

	//second
	if strings.ContainsAny(format, "s") {
		var second string
		if t.Second() < 10 {
			second = "0" + strconv.Itoa(t.Second())
		} else {
			second = strconv.Itoa(t.Second())
		}
		if strings.Count(format, "ss") == 1 && strings.Count(format, "s") == 2 {
			format = strings.Replace(format, "ss", second, 1)
		} else {
			panic("format second error! please 'ss'")
		}
	}
	return format
}

func main() {
	g := gin.Default()
	g.LoadHTMLFiles("./register.html")
	g.POST("sms", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", nil)
	})
	httpPostForm()
}

//GenerateNumberCode ...
func GenerateNumberCode(length int) string {
	code := make([]string, length)
	for i := 0; i < length; i++ {
		code[i] = strconv.Itoa(rand.Intn(10))
	}
	return strings.Join(code, "")
}
