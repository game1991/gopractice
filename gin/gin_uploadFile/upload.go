package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//api参数

	//限制上传文件大小为不大于8M,默认32M
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		//表单取文件
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"state": "400",
				"err":   err.Error(),
			})
		}
		//传到项目根目录,名字使用本身
		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file failed,err:%s\n", err.Error()))
				return
			}
			c.String(200, fmt.Sprintf("upload file [%s] success!\n", file.Filename))
		}

		//打印信息
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	router.Run(":8080")
}
