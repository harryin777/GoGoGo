package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128985530
// 数组合并
func main9() {
	var batch, count int
	fmt.Scan(&batch, &count)
	lists := make([][]string, 0, count)
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		str := buf.Text()
		if len(str) == 0 {
			continue
		}
		lists = append(lists, strings.Split(str, ","))
	}
	//fmt.Println(lists)
	cal9(batch, lists)

}

func cal9(fixed int, lists [][]string) {
	ans := make([]string, 0, len(lists))
	for i := 0; i < len(lists); i++ {
		if len(lists[i]) == 0 {
			continue
		}
		if len(lists[i]) >= fixed {
			for j := 0; j < fixed; j++ {
				ans = append(ans, lists[i][0])
				lists[i] = lists[i][1:]
			}
		} else {
			for len(lists[i]) > 0 {
				ans = append(ans, lists[i][0])
				lists[i] = lists[i][1:]
			}
		}
		if i == len(lists)-1 {
			i = -1
		}
	}
	res := ""
	for i := 0; i < len(ans); i++ {
		res = res + fmt.Sprintf("%v,", ans[i])
	}

	fmt.Println(res[:len(res)-1])
}
