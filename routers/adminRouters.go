package routers

import (
	"gindemo02/controllers/admin"

	"github.com/gin-gonic/gin"
)

type Admin struct {
	Account  string
	Password string
}

func adminMiddleware(c *gin.Context) {
	// c.Query("account")
	c.Set("MiddKey", "MiddVal")
}

func AdminRouterInit(r *gin.Engine) {

	adminRouter := r.Group("/admin")
	adminRouter.Use(adminMiddleware)
	{

		adminRouter.POST("/add", admin.UserConrterller{}.Add)
		adminRouter.PUT("/edit", admin.UserConrterller{}.Edit)
		adminRouter.GET("/infov2", admin.UserConrterller{}.Info)
		adminRouter.GET("/index", admin.UserConrterller{}.Index)
		adminRouter.GET("/getUsers", admin.UserConrterller{}.GetUsers)
		adminRouter.POST("/addV2", admin.UserConrterller{}.AddV2)
		adminRouter.DELETE("/delete", admin.UserConrterller{}.Delete)

		adminRouter.GET("/info", func(c *gin.Context) {
			account := c.Query("account")
			password := c.Query("password")
			admin := &Admin{
				Account:  account,
				Password: password,
			}
			c.JSON(200, admin)
		})

		adminRouter.POST("/login", func(c *gin.Context) {
			admin := &Admin{}
			c.BindJSON(admin)
			if admin.Account != "admin" || admin.Password != "123456" {
				c.JSON(401, gin.H{
					"message": "Login failed",
				})
				return
			} else {
				c.JSON(200, gin.H{
					"message": "Login success",
				})
			}
		})
	}

}
