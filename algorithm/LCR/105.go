package LCR

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	visitedFormat := "%v_%v"
	directionArr := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var dfs func(x, y int, visited map[string]struct{}) int
	dfs = func(x, y int, visited map[string]struct{}) int {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return 0
		}
		if grid[x][y] == 0 {
			return 0
		}
		key := fmt.Sprintf(visitedFormat, x, y)
		if _, ok := visited[key]; ok {
			return 0
		}
		res := 0
		visited[key] = struct{}{}
		for _, val := range directionArr {
			res += dfs(x+val[0], y+val[1], visited)
		}

		return res + 1
	}

	visited := make(map[string]struct{})
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			curr := dfs(i, j, visited)
			result = max(result, curr)
		}
	}

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
