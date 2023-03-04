package main

import "fmt"

// https://dream.blog.csdn.net/article/details/128986089
// 机器人走迷宫
func main27() {
	var col, row int
	fmt.Scan(&col, &row)
	var count int
	fmt.Scan(&count)
	wall := make([][]int, 0, count)
	for i := 0; i < count; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		wall = append(wall, []int{x, y})
	}

	cal26(col, row, wall)
}

func cal26(col, row int, wall [][]int) {
	grid := make([][]int, row)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, col)
	}
	for _, wal := range wall {
		grid[wal[1]][wal[0]] = -1
	}

	succList := make([][]int, 0, col*row)
	failList := make([][]int, 0, col*row)
	var dfs func(int, int) bool
	dfs = func(y, x int) bool {
		if x >= col || y >= row || grid[y][x] == -1 {
			return false
		}
		grid[y][x] = 2
		if x == col-1 && y == row-1 {
			return true
		}

		flag := false
		if dfs(y+1, x) {
			flag = true
			succList = append(succList, []int{y + 1, x})
		}
		if dfs(y, x+1) {
			flag = true
			succList = append(succList, []int{y, x})
		}
		if flag {
			return true
		}

		failList = append(failList, []int{y, x})
		return false
	}
	dfs(0, 0)
	empty := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				empty++
			}
		}
	}
	fmt.Printf("%d %d", len(failList), empty)
}
