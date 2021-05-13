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
