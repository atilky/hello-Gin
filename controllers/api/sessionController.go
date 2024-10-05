package api

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionDemo struct {
	Username string
}

type SessionController struct {
}

func (s SessionController) SetSession(c *gin.Context) {

	sessionDemo := &SessionDemo{}
	c.BindJSON(sessionDemo)
	fmt.Println(sessionDemo.Username)

	session := sessions.Default(c)
	session.Set("username", sessionDemo.Username)
	session.Save()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func (s SessionController) GetSession(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	c.JSON(200, gin.H{
		"code":     200,
		"username": username,
	})
}
