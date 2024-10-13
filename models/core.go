package models

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"gindemo02/util"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB
var dblog ormlog.Interface

var Err error

func InitMySql() {

	rootPath := util.ProjectRootPath
	config, err := ini.Load(rootPath + "/conf/app.ini")
	if err != nil {
		fmt.Println("fail to read file: %v", err)
		os.Exit(1)
	}

	//dblog := ormlog.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	ormlog.Config{
	//		SlowThreshold: 100 * time.Microsecond, // slow sql
	//		LogLevel:      ormlog.Info,            // Log level
	//		Colorful:      true,                   // 禁用彩色打印
	//	},
	//)

	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	dbname := config.Section("mysql").Key("database").String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, ip, port, dbname)

	fmt.Println("dsn : ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblog, PrepareStmt: true})
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblog, PrepareStmt: true}) //启用PrepareStmt，SQL预编译，提高查询效率

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	util.LogRus.Infof("connect to mysql db %s", dbname)

	DB = db
	Err = err
}
