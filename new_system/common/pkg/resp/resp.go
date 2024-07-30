package resp

import "github.com/gin-gonic/gin"

func Return(ctx *gin.Context, code int, msg string, data interface{}) {
	status := "success"
	if code != SuccessCode {
		status = "failure"
	}

	g := gin.H{
		"code":   code,
		"msg":    msg,
		"status": status,
		"data":   data,
	}

	ctx.JSON(200, g)
}
