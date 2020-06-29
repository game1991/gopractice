package main

import (
	"fmt"
	"net/http"
	"test/route"

	"html/template"

	_ "test/docs"

	"github.com/gin-gonic/gin"
)

//Data 数据
type Data struct {
	UserName string
	Code     string
}

//APIError ...
type APIError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// @Summary 展示模板
// @Tags template
// @Description a template
// @Produce  json
// @Success 200 string string "ok"
// @Failure 400 object APIError "err"
// @Router /template/ [get]
func templateFunc() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 解析指定文件生成模板对象
		tmpl, err := template.ParseFiles("./emailbind.tpl")
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}
		data := Data{
			UserName: "ganlei",
			Code:     "875466",
		}
		// buf := bytes.NewBuffer(nil)

		// // 利用给定数据渲染模板，并将结果写入w
		// tmpl.Execute(buf, data)
		// fmt.Printf("result is %#v\n", string(buf.Bytes()))
		if err := tmpl.Execute(c.Writer, data); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, APIError{Code: 400, Msg: err.Error()})
			return
		}
	}
}

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name ganlei
// @contact.url https://juejin.im/user/5e9f9b15f265da47b55504de
// @contact.email ganlei@uniontech.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
func main() {

	engine := gin.Default()
	route.InitRouter(engine)
	r := engine.Group("template")
	r.GET("", templateFunc())
	if err := engine.Run(":9090"); err != nil {
		panic(err)
	}
}
