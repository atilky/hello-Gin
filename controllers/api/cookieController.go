package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CookieController struct {
}

func (c CookieController) SetCookie(g *gin.Context) {
	g.SetCookie("gin_cookie", "test_cookie", 3600, "/", "localhost", false, true)
	g.String(200, "設置 cookie 成功")
}

func (c CookieController) GetCookie(g *gin.Context) {
	cookie, err := g.Cookie("gin_cookie")
	if err != nil {
		g.JSON(200, gin.H{
			"code": 200,
			"msg":  "get cookie error, " + err.Error(),
		})
		return
	}
	fmt.Println("cookie:", cookie)
	g.JSON(200, gin.H{
		"code": 200,
		"msg":  "get cookie success, " + cookie,
	})
}
