package routers

import (
	"gindemo02/controllers/admin"
	"github.com/gin-gonic/gin"
)

func LoginRouterInit(g *gin.Engine) {

	g.POST("/login", admin.LoginController{}.Login)
	g.GET("/getauthtoken", admin.LoginController{}.GetAuthToken)

}
