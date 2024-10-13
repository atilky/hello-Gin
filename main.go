package main

import (
	"fmt"
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

	r := gin.Default()
	// r.GET("/",
	// 	func(c *gin.Context) {
	// 		fmt.Println("Hello middleware")
	// 	},
	// 	func(c *gin.Context) {
	// 		c.String(200, "gin 首頁")
	// 	},
	// )

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
	setSessionByRedis(r)

	// 演示 gopkg.in/ini.v1
	// config, err := ini.Load("./conf/app.ini")
	// if err != nil {
	// 	fmt.Println("fail to read file: %v", err)
	// 	os.Exit(1)
	// }

	// get ini value
	// fmt.Println("app_name:", config.Section("").Key("app_name").String())
	// fmt.Println("app_mode:", config.Section("mysql").Key("password").String())

	// set ini value
	// config.Section("").Key("app_name").SetValue("production")
	// config.SaveTo("./conf/app.ini")

	rootPath := util.ProjectRootPath
	r.Static("/static", rootPath+"/conf")

	routers.AdminRouterInit(r)
	routers.ApiRouterInit(r)
	routers.RestfulDemo(r)
	routers.HealthCheck(r)

	r.Run(":80")

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
