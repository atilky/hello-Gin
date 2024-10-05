package models

import "time"

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetNowTime 获取当前时间 格式：2006-01-02 15:04:05
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetNowTimeUnix 获取当前时间戳
func GetNowTimeUnix() int64 {
	return time.Now().Unix()
}

// GetNowDay 获取当前日期 格式：20060102
func GetNowDay() string {
	return time.Now().Format("20060102")
}
