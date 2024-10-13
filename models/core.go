package models

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"gindemo02/util"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Err error

func init() {

	rootPath := util.ProjectRootPath

	config, err := ini.Load(rootPath + "/conf/app.ini")
	if err != nil {
		fmt.Println("fail to read file: %v", err)
		os.Exit(1)
	}

	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	dbname := config.Section("mysql").Key("database").String()
	// allowPublicKeyRetrieval := config.Section("mysql").Key("allowPublicKeyRetrieval").String()

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, ip, port, dbname)

	fmt.Println("dsn : ", dsn)

	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
