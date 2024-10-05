package admin

import (
	"fmt"
	"gindemo02/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

type UserConrterller struct {
	BaseController
}

func (u UserConrterller) Add(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	c.JSON(200, user)
}

func (u UserConrterller) Edit(c *gin.Context) {
	// c.JSON(200, "edit ~")
	successStr := c.DefaultQuery("isSuccess", "false")
	success, err := strconv.ParseBool(successStr)

	val, _ := c.Get("MiddKey")
	fmt.Println(val)

	if err != nil {
		u.Error(c, 400, "edit failed")
		return
	}

	if success {
		u.Success(c, "edit success")
	} else {
		u.Error(c, 400, "edit failed")
	}
}

func (u UserConrterller) Info(c *gin.Context) {

	name := c.Query("name")
	fmt.Println(name)

	switch name {
	case "zhangsan":
		c.JSON(200, User{
			Name: "zhangsan",
			Age:  18,
		})

	case "alan":
		c.JSON(200, User{
			Name: "alan",
			Age:  20,
		})

	default:
		c.JSON(200, gin.H{
			"message": "User not found",
		})
	}
}

func (u UserConrterller) Index(c *gin.Context) {

	users := []models.User{}
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": users,
	})

}

func (u UserConrterller) AddV2(c *gin.Context) {

	users := []models.User{
		{Username: "russ", Age: 38, Email: "s2@test.com", AddTime: models.GetNowTimeUnix()},
		{Username: "colin", Age: 28, Email: "s3@test.com", AddTime: models.GetNowTimeUnix()},
	}

	result := models.DB.Create(users)
	if result.Error != nil {
		u.Error(c, 400, "add failed")
		return
	}

	u.Success(c, "add success")

}
