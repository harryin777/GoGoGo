package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128989841
// 任务执行总时长
func main16() {
	var aT, bT, num int
	var str string
	fmt.Scan(&str)
	arr := strings.Split(str, ",")
	for i := 0; i < len(arr); i++ {
		tmp, _ := strconv.Atoi(arr[i])
		if i == 0 {
			aT = tmp
		} else if i == 1 {
			bT = tmp
		} else {
			num = tmp
		}
	}

	cal16(aT, bT, num)
}

func cal16(a, b, num int) {
	data := []int{a, b}
	total := make([]int, 0, num)
	dup := make(map[int]struct{})
	var dfs func(int, int)
	dfs = func(cur int, count int) {
		if count == num {
			if _, e := dup[cur]; !e {
				dup[cur] = struct{}{}
				total = append(total, cur)
			}
			return
		}

		for i := 0; i < len(data); i++ {
			cur += data[i]
			dfs(cur, count+1)
			cur -= data[i]
		}
	}
	dfs(0, 0)
	sort.Slice(total, func(i, j int) bool {
		return total[i] < total[j]
	})
	fmt.Println(total)
}
