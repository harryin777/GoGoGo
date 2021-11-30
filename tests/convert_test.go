/**
 * @Author: yky
 * @Date: 2021/5/13 12:21
 * @Version: 1.0
 */
package tests

import (
	"fmt"
	"test1/Utils"
	"testing"
	"time"
)

const (
	AlbumListKeyByPageKey = "albumList:%v"
)

func TestTime(t *testing.T) {
	param := "2014-10-27T16:00:00.000+00:00"
	_ = param
	Utils.Time()

	page := "1"
	fmt.Println(fmt.Sprintf(AlbumListKeyByPageKey, page))
}

func Test_DateTOStr(t *testing.T) {
	fmt.Println(Utils.DateTOStr(time.Now()))
}

func Test_GetTimeArr(t *testing.T) {
	dateStr1 := "2021-05-19"
	dateStr2 := "2021-05-18"
	fmt.Println(Utils.GetInterval(dateStr2, dateStr1))
}
