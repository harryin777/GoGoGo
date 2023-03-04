package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/129067043
// 吃火锅
func main50() {
	var dishCount, sec int
	fmt.Scan(&dishCount, &sec)
	data := make([][]int, 0, dishCount)
	for i := 0; i < dishCount; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		data = append(data, []int{start, start + end})
	}

	cal50(data, sec)
}

func cal50(data [][]int, sec int) {
	sort.Slice(data, func(i, j int) bool {
		return data[i][1] < data[j][1]
	})

	ans := make([]int, len(data))
	ans[0] = 1
	pre := 0
	for i := 1; i < len(data); i++ {
		if data[i][1] >= data[pre][1]+sec {
			ans[i] = 1
			pre = i
		}
	}

	count := 0
	for i := 0; i < len(ans); i++ {
		if ans[i] == 1 {
			count++
		}
	}

	fmt.Println(count)
}
