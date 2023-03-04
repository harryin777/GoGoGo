package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129023694
// 聚餐地点
func main40() {
	var row, col int
	fmt.Scan(&row, &col)
	grid := make([][]cordinate, 0, row)
	var p1, p2 []int
	for i := 0; i < row; i++ {
		arr := make([]cordinate, 0, col)
		for j := 0; j < col; j++ {
			var t int
			fmt.Scan(&t)
			if t == 2 {
				if len(p1) == 0 {
					p1 = []int{i, j}
				} else {
					p2 = []int{i, j}
				}
			}
			arr = append(arr, cordinate{
				X:      i,
				Y:      j,
				Status: t,
			})
		}
		grid = append(grid, arr)
	}

	cal40(grid, p1, p2)
}

func cal40(grid [][]cordinate, p1, p2 []int) {
	routeMap := make(map[cordinate]int)
	routeMap2 := make(map[cordinate]int)
	var dfs func(int, int, int, [][]cordinate, int)
	dfs = func(x int, y int, pace int, grid [][]cordinate, who int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y].Status == 1 || grid[x][y].Status == 4 {
			return
		}

		if grid[x][y].Status == 3 {
			if who == 1 {
				c := cordinate{
					X: x,
					Y: y,
				}
				if routeMap[c] != 0 {
					pace = min(routeMap[c], pace)
				}
				routeMap[c] = pace
			} else {
				c := cordinate{
					X: x,
					Y: y,
				}
				if routeMap2[c] != 0 {
					pace = min(routeMap2[c], pace)
				}
				routeMap2[c] = pace
			}

			return
		}
		grid[x][y].Status = 4
		dfs(x+1, y, pace+1, grid, who)
		dfs(x, y+1, pace+1, grid, who)
		dfs(x-1, y, pace+1, grid, who)
		dfs(x, y-1, pace+1, grid, who)
		grid[x][y].Status = 0
	}
	g1 := make([][]cordinate, len(grid))
	g2 := make([][]cordinate, len(grid))
	for i := 0; i < len(grid); i++ {
		g1[i] = make([]cordinate, len(grid[0]))
		g2[i] = make([]cordinate, len(grid[0]))
	}
	copy(g1, grid)
	copy(g2, grid)
	dfs(p1[0], p1[1], 0, g1, 1)
	dfs(p2[0], p2[1], 0, g2, 2)
	fmt.Println()

}

type cordinate struct {
	X      int
	Y      int
	Status int
}
