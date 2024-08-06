package resp

import "github.com/gin-gonic/gin"

type RespData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Return(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, RespData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
