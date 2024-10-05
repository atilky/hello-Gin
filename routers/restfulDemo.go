package routers

import (
	"github.com/gin-gonic/gin"
)

type Article struct {
	Tilte string
	Desc  string
}

func RestfulDemo(r *gin.Engine) {

	// router demo
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"message": "Hello, JSON1",
			"success": true,
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		article := Article{
			Tilte: "Hello, JSON2",
			Desc:  "This is a description",
		}
		c.JSON(200, article)
	})

	r.GET("/json3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, JSON3",
			"success": true,
		})
	})

	r.GET("/json4", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "1")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// post x-www-form-urlencoded
	r.POST("/json5", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "1")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// request body
	r.POST("/json6", func(c *gin.Context) {
		article := &Article{}
		c.BindJSON(article)
		c.JSON(200, gin.H{
			"Desc":  article.Desc,
			"Tilte": article.Tilte,
		})
	})

}
