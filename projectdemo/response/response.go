package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//用于统一回复格式

/*
格式如下
{
	code:200001,
	data:xxx,
	msg:xx
}*/

func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
