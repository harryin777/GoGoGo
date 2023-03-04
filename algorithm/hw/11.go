package main

import (
	"fmt"
	"sort"
)

func main11() {
	var m, n, R int
	fmt.Scan(&m, &n, &R)
	A := make([]int, 0, m)
	B := make([]int, 0, m)
	for i := 0; i < m; i++ {
		var tmp int
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		B = append(B, tmp)
	}

	cal11(A, B, R)
}

func cal11(a, b []int, r int) {
	ans := make([][][]int, len(a))
	for i := 0; i < len(a); i++ {
		ans[i] = make([][]int, 0, len(b))
		for j := 0; j < len(b); j++ {
			if a[i] <= b[j] && b[j]-a[i] <= r {
				ans[i] = append(ans[i], []int{a[i], b[j]})
			}
		}
	}

	for i := 0; i < len(ans); i++ {
		if len(ans[i]) == 0 {
			continue
		}
		sort.Slice(ans[i], func(m, n int) bool {
			return ans[i][m][0] < ans[i][n][0]
		})
		fmt.Printf("%v %v \n", ans[i][0][0], ans[i][0][1])
	}
}
