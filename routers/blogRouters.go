package routers

import (
	"gindemo02/controllers/blog"

	"github.com/gin-gonic/gin"
)

type blogRouters struct {
}

func BlogRouterInit(r *gin.Engine) {

	blogRouter := r.Group("/blog")
	//adminRouter.Use(adminMiddleware)
	{
		blogRouter.PUT("/edit", blog.BlogController{}.Edit)
	}

}
