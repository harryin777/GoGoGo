package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/129103167
// 叠放书籍
func main45() {
	data := make([][]int, 0, 10)
	var s string
	fmt.Scan(&s)
	reg := regexp.MustCompile("\\d{1,9},\\d{2}")
	arr := reg.FindAllStringSubmatch(s, -1)
	for i := 0; i < len(arr); i++ {
		tmp := strings.Split(arr[i][0], ",")
		t1, _ := strconv.Atoi(tmp[0])
		t2, _ := strconv.Atoi(tmp[1])
		data = append(data, []int{t1, t2})
	}
	cal45(data)
}

func cal45(data [][]int) {
	sort.Slice(data, func(i, j int) bool {
		if data[i][0] < data[j][0] {
			return true
		} else if data[i][0] == data[j][0] {
			return data[i][1] < data[j][1]
		} else {
			return false
		}
	})
	dp := make([]int, len(data))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	//ans := ^int(^uint(0) >> 1)
	ans := 0
	for i := 1; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[i][0] > data[j][0] && data[i][1] > data[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}

	fmt.Println(ans)

}
