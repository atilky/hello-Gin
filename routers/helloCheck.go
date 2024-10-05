package routers

import (
	"github.com/gin-gonic/gin"
)

func HealthCheck(r *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": 200,
		})
	})

}
