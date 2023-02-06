/**
 * @Author: yky
 * @Date: 2021/5/13 12:21
 * @Version: 1.0
 */
package Utils

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func Time() {
	fmt.Println("time ===============:")
	//2006-01-02 15:04:05 时间必须是这个
	const base_format = "2006-01-02"
	//获取当前时间
	nt := time.Now()
	fmt.Printf("now datetime:%v\n", nt)

	//延时年月日
	adnt := nt.AddDate(1, 2, 3)
	fmt.Printf("now adddate:%v\n", adnt)

	//延时秒
	ant := nt.Add(3600 * 1e9) //延时1小时=60*60 秒
	fmt.Printf("now add:%v\n", ant)

	//转换为时间格式字符串
	str_time := nt.Format(base_format)
	fmt.Printf("now time string:%v\n", str_time)

	//时间字符串转换为日期格式
	parse_str_time, _ := time.Parse(base_format, str_time)
	fmt.Printf("string to datetime :%v\n", parse_str_time)

	//时间戳 秒
	timestamp := time.Now().Unix()
	println("timestamp:", timestamp)
	//时间戳 毫秒
	msec := time.Now().UnixNano() / 1e6
	println("timestamp msec:", msec)

	float_ms := msec % timestamp
	v := fmt.Sprintf("%.3f\n", float64(float_ms)/1000.0)
	println("float msec:", v)

	//时间戳转日期格式
	date_time := time.Unix(timestamp, 0)
	fmt.Printf("timestamp to datetime:%v\n", date_time)

	//时间字符串转时间戳
	t, _ := time.Parse(base_format, "2014-10-27T16:00:00.000+00:00")
	datetime_str_to_timestamp := t.Unix()
	println("datetime_str_to_timestamp:", datetime_str_to_timestamp)

}

func DateTOStr(currentTime time.Time) string {
	const dateFmt = "2006-01-02"
	str_time := currentTime.Format(dateFmt)
	return str_time
}

// 安全性会出问题,defer + recover 都不会捕获,因为 string 底层是不可变的, byte 数组是可以对数组内元素做变更的. 只有在只读的情况下可用
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
