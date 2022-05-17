/**
 * @Author: yky
 * @Date: 2021/5/19 17:08
 * @Version: 1.0
 */
package Utils

import "time"

const (
	FullLayout = "2006-01-02 15:04:05"
)

func GetInterval(start, end string) int64 {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout, start, loc)
	endUnix, _ := time.ParseInLocation(timeLayout, end, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	days := (endTime - startTime) / 86400
	return days
}
