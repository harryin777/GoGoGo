package main

import (
	"fmt"
	"strconv"
)

// https://dream.blog.csdn.net/article/details/129067111
// 数字的排列
func main6() {
	var n int
	fmt.Scan(&n)

	cal6(n)
}

func cal6(n int) {
	list := make([][]int, n)
	for i := 0; i < len(list); i++ {
		list[i] = make([]int, 0, 0)
	}
	point := 1
	for i := 0; i < len(list); i++ {
		if (i+1)%2 == 1 {
			for j := 0; j < i+1; j++ {
				list[i] = append(list[i], point)
				point++
			}
		} else {
			for j := 0; j < i+1; j++ {
				list[i] = append([]int{point}, list[i]...)
				point++
			}
		}
	}

	finalTri := make([][]string, n)
	for i := 0; i < len(list); i++ {
		finalTri[i] = make([]string, 0, 100)
		for j := 0; j < len(list[i]); j++ {
			numStr := strconv.FormatInt(int64(list[i][j]), 10)
			for len(numStr) < 4 {
				numStr += "*"
			}
			if j == 0 {
				for m := 0; m < n-i-1; m++ {
					numStr = "    " + numStr
				}
			} else {
				numStr = "    " + numStr
			}
			finalTri[i] = append(finalTri[i], numStr)
		}
	}

	for i := 0; i < len(finalTri); i++ {
		fmt.Println(finalTri[i])
	}
}
