package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/129095347
// 统计匹配的二元组个数
func main46() {
	var m, n int
	fmt.Scan(&m, &n)
	data1 := make([]int, 0, m)
	data2 := make([]int, 0, n)
	//str1, str2 := "", ""
	for i := 0; i < m; i++ {
		var t int
		fmt.Scan(&t)
		//str1 += strconv.FormatInt(int64(t), 10)
		data1 = append(data1, t)
	}
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		//str2 += strconv.FormatInt(int64(t), 10)
		data2 = append(data2, t)
	}
	cal46(data1, data2)
}

func cal46(data1 []int, data2 []int) {
	sort.Slice(data1, func(i, j int) bool {
		return data1[i] < data1[j]
	})

	sort.Slice(data2, func(i, j int) bool {
		return data2[i] < data2[j]
	})

	count := 0
	if len(data1) > len(data2) {
		for i := 0; i < len(data1); i++ {
			if index := bs(0, len(data2), data1[i], data2); index != -1 && index < len(data2) {
				count++
			}
		}
	} else {
		for i := 0; i < len(data2); i++ {
			if index := bs(0, len(data1), data2[i], data1); index != -1 && index < len(data1) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func bs(left, right, target int, data []int) int {
	if left > right {
		return -1
	}

	for left < right {
		mid := (left + right) >> 1
		if target > data[mid] {
			left = mid + 1
		} else if target < data[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return left
}
