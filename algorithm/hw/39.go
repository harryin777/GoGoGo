package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129019205
// 基站维修工程师
func main39() {
	var r int
	fmt.Scan(&r)
	grid := make([][]int, 0, r)
	for i := 0; i < r; i++ {
		tmp := make([]int, 0, r)
		for j := 0; j < r; j++ {
			var t int
			fmt.Scan(&t)
			tmp = append(tmp, t)
		}
		grid = append(grid, tmp)
	}

	cal39(grid)
}

func cal39(grid [][]int) {
	short := int(^uint(0) >> 1)
	var dfs func(int, int, int)
	dfs = func(pace int, currStation int, route int) {
		if pace > len(grid) || (pace == len(grid) && currStation != 0) {
			return
		}

		if pace == len(grid) && currStation == 0 {
			short = min(short, route)
			return
		}

		for i := 0; i < len(grid[currStation]); i++ {
			if i == currStation || (pace != len(grid)-1 && i == 0) {
				continue
			}
			dfs(pace+1, i, route+grid[currStation][i])
		}
	}
	dfs(0, 0, 0)

	fmt.Println(short)
}

//func cal39(grid [][]int) {
//	short := int(^uint(0) >> 1)
//	for i := 0; i < len(grid); i++ {
//		if i == 0 {
//			continue
//		}
//		cur := grid[i][0] + grid[0][i]
//		short = min(short, cur)
//	}
//
//	fmt.Println(short)
//}
