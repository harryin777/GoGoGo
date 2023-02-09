/**
*   @Author: yky
*   @File: http_test
*   @Version: 1.0
*   @Date: 2021-06-13 10:34
 */
package tests

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

var url []string

func init() {
	url = []string{
		"https://www.baidu.com/",
	}
}

/*
*
  - @Description: io.copy防止内容全部读取到内存造成内存溢出
    http请求
  - @Param:
  - @return:

*
*/
func Test_HttpSimple(t *testing.T) {
	url := "http://wx.qlogo.cn/Vaz7vE1/64"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}

	defer resp.Body.Close()

	out, err := os.Create("/tmp/icon_wx_2.png")
	wt := bufio.NewWriter(out)

	//defer out.Close()

	n, err := io.Copy(wt, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		panic(err)
	}
	wt.Flush()
}
