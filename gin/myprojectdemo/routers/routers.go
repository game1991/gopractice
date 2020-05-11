package routers

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

//注册App的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.Default()
	for _, opt := range options {
		opt(r)
	}
	return r
}
