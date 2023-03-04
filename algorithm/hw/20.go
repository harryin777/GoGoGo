package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/128989935
// 整数分解

func main20() {
	var count int
	fmt.Scan(&count)
	cal20(count)
}

func cal20(n int) {
	ans := make([][]int, 0, n)
	for i := 1; i <= n; i++ {
		tmp := make([]int, 0, 20)
		cur := 0
		flag := false
		for j := i; j <= n; j++ {
			tmp = append(tmp, j)
			cur += j
			if cur == n {
				flag = true
				break
			} else if cur > n {
				break
			}
		}
		if flag {
			ans = append(ans, tmp)
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return len(ans[i]) < len(ans[j])
	})
	for i := 0; i < len(ans); i++ {
		scr := fmt.Sprintf("%d=", n)
		for j := 0; j < len(ans[i]); j++ {
			scr += fmt.Sprintf("%d+", ans[i][j])
		}
		scr = scr[:len(scr)-1]
		fmt.Println(scr)
	}
	fmt.Printf("Result:%d", len(ans))
}
