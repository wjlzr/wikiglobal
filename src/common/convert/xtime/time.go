package xtime

import "time"

func DateFormat(time1 string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", time1, loc)
	return tt.Unix()
}
