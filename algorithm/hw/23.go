package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128989994
// 子序列长度
func main23() {
	var arr []int
	var num int
	var str string
	fmt.Scan(&str)
	strs := strings.Split(str, ",")
	for _, i2 := range strs {
		tmp, _ := strconv.Atoi(i2)
		arr = append(arr, tmp)
	}
	fmt.Scan(&num)
	cal23(arr, num)
}

func cal23(arr []int, num int) {
	ans := make([][]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := i; j < len(arr); j++ {
			count += arr[j]
			if count == num {
				ans = append(ans, arr[i:j+1])
			} else if count > num {
				break
			}
		}
	}

	if len(ans) == 0 {
		fmt.Println(-1)
		return
	}

	sort.Slice(ans, func(i, j int) bool {
		return len(ans[i]) > len(ans[j])
	})
	fmt.Println(len(ans[0]))
}
