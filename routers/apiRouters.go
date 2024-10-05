package routers

import (
	"gindemo02/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(g *gin.Engine) {

	apiRouter := g.Group("/api")
	{
		apiRouter.GET("/timestamp", api.ToolsController{}.GetTimeStr)
		apiRouter.POST("/setCookie", api.CookieController{}.SetCookie)
		apiRouter.GET("/getCookie", api.CookieController{}.GetCookie)
		apiRouter.POST("/setSession", api.SessionController{}.SetSession)
		apiRouter.GET("/getSession", api.SessionController{}.GetSession)

		apiRouter.GET("/json1", func(c *gin.Context) {
			c.JSON(200, map[string]interface{}{
				"message": "Hello, api JSON1",
				"success": true,
			})
		})

	}
}
