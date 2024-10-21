package models

import (
	"fmt"
	"gindemo02/util"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"os"
)

var (
	REDIS *redis.Client
)

func createRedisClient(address, passwd string, db int) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: passwd,
		DB:       db,
	})
	if err := cli.Ping().Err(); err != nil {
		util.LogRus.Panicf("connect to redis %d failed %v", db, err)
	} else {
		util.LogRus.Infof("connect to redis %d", db) //能ping成功才说明连接成功
	}
	return cli
}

func InitRedis() {

	rootPath := util.ProjectRootPath
	config, err := ini.Load(rootPath + "/conf/app.ini")
	if err != nil {
		fmt.Println("fail to read file: %v", err)
		os.Exit(1)
	}

	addr := config.Section("redis").Key("ip").String()
	pass := config.Section("redis").Key("pass").String()
	db, _ := config.Section("redis").Key("db").Int()

	REDIS = createRedisClient(addr, pass, db)

}
