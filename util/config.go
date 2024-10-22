package util

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path"
	"runtime"
)

var (
	// 取得根路徑
	ProjectRootPath = path.Dir(getoncurrentPath()+"/../") + "/"
	ConfigMap       = make(map[string]map[string]string)
)

func getoncurrentPath() string {
	_, filename, _, _ := runtime.Caller(0) //0表示当前本行代码在什么位置
	return path.Dir(filename)              //返回文件所在的目录
}

func init() {
	rootPath := ProjectRootPath
	config, err := ini.Load(rootPath + "/conf/app.ini")
	if err != nil {
		fmt.Println("fail to read file: %v", err)
		os.Exit(1)
	}

	ConfigMap["key"] = make(map[string]string)
	ConfigMap["key"]["jwtKey"] = config.Section("key").Key("jwtKey").String()

}
