package main

import (
	"fmt"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128995153
// 最大相连男生数
func main33() {
	var row, column int
	fmt.Scanf("%d,%d", &row, &column)
	grid := make([][]string, 0, row)
	for i := 0; i < row; i++ {
		var tmp string
		fmt.Scan(&tmp)
		grid = append(grid, strings.Split(tmp, ","))
	}

	cal33(grid)
}

func cal33(grid [][]string) {
	actions := [][]int{
		{1, 0},
		{0, 1},
		{1, 1},
		{-1, -1},
	}
	maxLen := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for m := 0; m < len(actions); m++ {
				count := 0
				r, c := i, j
				for r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
					if grid[r][c] == "M" {
						count++
					} else {
						maxLen = max(maxLen, count)
						count = 0
					}
					r += actions[m][0]
					c += actions[m][1]
				}
				maxLen = max(maxLen, count)
			}
		}
	}

	fmt.Println(maxLen)
}
