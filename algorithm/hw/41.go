package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/129023719
// 不含101
func main41() {
	var begin, end int
	fmt.Scan(&begin, &end)
	cal41(begin, end)
}

func cal41(begin, end int) {
	ans := 0
	for i := begin; i < end+1; i++ {
		if !strings.Contains(T2B(i), "101") {
			ans++
		}
	}
	fmt.Println(ans)
}

func T2B(val int) string {
	res := ""
	ans := make([]int, 0, val)
	for val != 0 {
		ans = append(ans, val%2)
		val /= 2
	}
	for i := len(ans) - 1; i >= 0; i-- {
		res += strconv.FormatInt(int64(ans[i]), 10)
	}
	return res
}
