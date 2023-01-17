package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespOK(c *gin.Context, info string) { //正确
	OK.info = info
	c.JSON(http.StatusOK, OK)
}
func RespParamErr(c *gin.Context) { //错误
	c.JSON(http.StatusBadRequest, ParamError)
}

// 服务器内部错误
func RsepInternalErr(c *gin.Context) { //连接错误
	c.JSON(http.StatusInternalServerError, InternalError)
}

func NormalErr(c *gin.Context, status int, info string) { //其他错误
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}

func RespSuccess(c *gin.Context, code int, mess string, data interface{}) {
	c.JSON(http.StatusOK, &Resp{
		Code:    code,
		Message: mess,
		Data:    data,
	})
}
