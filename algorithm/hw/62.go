package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129088380
// 查找单入口空闲区域
func main62() {
	var row, col int
	fmt.Scan(&row, &col)
	grid := make([][]string, 0, row)
	for i := 0; i < row; i++ {
		var tmp []string
		for j := 0; j < col; j++ {
			var node string
			fmt.Scan(&node)
			tmp = append(tmp, node)
		}
		grid = append(grid, tmp)
	}

	cal62(grid)
}

func cal62(grid [][]string) {
	//actions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(int, int, [][]string) bool
	var count int
	dfs = func(x int, y int, grid [][]string) bool {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == "X" {
			return false
		}
		grid[x][y] = "X"
		count++

		dfs(x+1, y, grid)
		dfs(x, y+1, grid)
		dfs(x-1, y, grid)
		dfs(x, y-1, grid)

		return false
	}

	var ans [][]int
	var mxOne int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[i])-1 {
				if grid[i][j] == "O" {
					tmp := make([][]string, len(grid), len(grid))
					for p := 0; p < len(tmp); p++ {
						tmp[p] = make([]string, len(grid[0]), len(grid[0]))
						copy(tmp[p], grid[p])
					}
					dfs(i, j, tmp)
					if mxOne == count {
						mxOne, count = 0, 0
						ans = [][]int{}
					} else if mxOne < count {
						mxOne = count
						count = 0
						ans = [][]int{{i, j}}
					}
				}
			}
		}
	}
	if len(ans) == 0 {
		fmt.Println("NULL")
	} else {
		fmt.Printf("%d %d %d", ans[0][0], ans[0][1], mxOne)
	}
}
