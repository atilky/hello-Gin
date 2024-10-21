package main

import (
	"fmt"
	"gindemo02/models"
	"gindemo02/routers"
	"gindemo02/util"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	beforeTime := time.Now().Unix()
	fmt.Println("Hello middleware handler before")
	c.Next()
	afterTime := time.Now().Unix()
	fmt.Println("Hello middleware handler after, time:", afterTime-beforeTime)

}

func main() {

	Init()

	r := gin.Default()

	r.GET("/",
		middleware,
		func(c *gin.Context) {
			fmt.Println("Hello gin 首頁")
			c.String(200, "gin 首頁")
		},
	)

	// set session
	// setSession(r)
	// set session by redis
	//setSessionByRedis(r)

	rootPath := util.ProjectRootPath
	r.Static("/static", rootPath+"/conf")

	routers.AdminRouterInit(r)
	routers.ApiRouterInit(r)
	routers.RestfulDemo(r)
	routers.HealthCheck(r)
	routers.BlogRouterInit(r)
	routers.LoginRouterInit(r)

	r.Run(":81")

}

func Init() {
	util.InitLog("log")
	models.InitMySql()
	models.InitRedis()
}

func setSession(r *gin.Engine) {

	// 創建基於 cookie 的存儲引擎，secret 用於加密的密鑰
	store := cookie.NewStore([]byte("secret"))
	// 配置 session 中間件, store 為上面創建的存儲引擎, 我們可以替換成其他
	r.Use(sessions.Sessions("mysession", store))
}

func setSessionByRedis(r *gin.Engine) {

	// 創建基於 cookie 的存儲引擎，secret 用於加密的密鑰
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// 配置 session 中間件, store 為上面創建的存儲引擎, 我們可以替換成其他
	r.Use(sessions.Sessions("mysession", store))
}
