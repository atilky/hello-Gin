package util

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"strings"
	"time"
)

var (
	LogRus *logrus.Logger
)

func InitLog(configFile string) {

	config, err := ini.Load(ProjectRootPath + "/conf/app.ini")
	if err != nil {
		fmt.Println("fail to read file: %v", err)
		os.Exit(1)
	}

	level := config.Section("log").Key("level").String()
	outFile := config.Section("log").Key("file").String()

	LogRus = logrus.New()
	switch strings.ToLower(level) {
	case "debug":
		LogRus.SetLevel(logrus.DebugLevel)
	case "info":
		LogRus.SetLevel(logrus.InfoLevel)
	case "warn":
		LogRus.SetLevel(logrus.WarnLevel)
	case "error":
		LogRus.SetLevel(logrus.ErrorLevel)
	case "panic":
		LogRus.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level %s", level))
	}

	LogRus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000", // 显示ms
	})
	logFile := ProjectRootPath + outFile
	fout, err := rotatelogs.New(
		logFile+".%Y%m%d%H%mmm",                  //指定日志文件的路径和名称，路径不存在时会创建
		rotatelogs.WithLinkName(logFile),         //为最新的一份日志创建软链接
		rotatelogs.WithRotationTime(1*time.Hour), //每隔1小时生成一份新的日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),    //只留最近7天的日志，或使用WithRotationCount只保留最近的几份日志
	)
	if err != nil {
		panic(err)
	}
	LogRus.SetOutput(fout) //设置日志文件
	//LogRus.SetReportCaller(true) //输出是从哪里调起的日志打印
}
