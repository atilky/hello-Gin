package util

import (
	"path"
	"runtime"
)

var (
	// 取得根路徑
	ProjectRootPath = path.Dir(getoncurrentPath()+"/../") + "/"
)

func getoncurrentPath() string {
	_, filename, _, _ := runtime.Caller(0) //0表示当前本行代码在什么位置
	return path.Dir(filename)              //返回文件所在的目录
}
