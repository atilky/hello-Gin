package admin

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (b BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func (b BaseController) Error(c *gin.Context, code int, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
	})
}
