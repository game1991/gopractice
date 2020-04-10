package main

import (
	"BMS/v1/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//BookMangeSystem
func main() {
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	engine := gin.Default()
	//模板渲染
	engine.LoadHTMLGlob("templates/**/*")

	//路由组
	g := engine.Group("book")
	g.GET("list", bookListHandler)
	g.GET("new", getBookHandler)
	g.POST("new", createBookHandler)
	g.GET("delete", deleteBookHandler)

	//运行
	engine.Run()
}

func bookListHandler(c *gin.Context) {
	bookList, err := db.QueryAllBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "1",
			"msg":  err.Error(),
		})
		return
	}
	//返回书籍信息
	/*
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"data":bookList,
		})
	*/
	c.HTML(http.StatusOK, "book/book_list.html", gin.H{
		"code": "0",
		"data": bookList,
	})

}

func getBookHandler(c *gin.Context) {
	//给用户返回一个添加书籍页面的处理函数
	c.HTML(http.StatusOK, "book/new_book.html", nil)
}

func createBookHandler(c *gin.Context) {
	//创建书籍的处理函数

	//从表单form中读取数据
	title := c.PostForm("title")
	price := c.PostForm("price")
	priceVal, err := strconv.ParseFloat(price, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  "无效的价格参数",
		})
		return
	}
	//插入数据库
	if err := db.InsertBook(title, priceVal); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  "插入书籍数据失败",
		})
		return
	}
	//插入数据成功
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

//删除书籍
func deleteBookHandler(c *gin.Context) {
	//取参数query string
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  "删除书籍数据错误",
		})
		return
	}
	//数据库去删除
	if err := db.DeleteBook(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  err.Error(),
		})
		return
	}
	//删除成功跳转书籍列表页
	c.Redirect(http.StatusMovedPermanently, "/book/list")

}
