package main

import (
	"fmt"
)

// https://dream.blog.csdn.net/article/details/128985593
// 计算机网络信号
func main10() {
	var m, n int
	fmt.Scan(&m, &n)
	grid := make([][]node, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]node, 0, n)
		for j := 0; j < n; j++ {
			var tmp int
			fmt.Scan(&tmp)
			grid[i] = append(grid[i], node{
				X:      i,
				Y:      j,
				Signal: tmp,
			})
		}
	}
	var x, y int
	fmt.Scan(&x, &y)
	//fmt.Println(grid)
	cal10(grid, x, y)
}

type node struct {
	X       int
	Y       int
	Signal  int
	Visited bool
}

func cal10(grid [][]node, x, y int) {
	tx, ty := 0, 0
	//signal := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].Signal > 0 {
				tx, ty = i, j
				//signal = grid[i][j].Signal
				break
			}
		}
	}

	queue := make([]*node, 0, 10)
	queue = append(queue, &grid[tx][ty])
	dirArr := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(queue) != 0 {
		ele := queue[0]
		queue = queue[1:]
		for _, val := range dirArr {
			ele.Visited = true
			newX, newY := ele.X+val[0], ele.Y+val[1]
			if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) {
				continue
			}
			if 0 == grid[newX][newY].Signal {
				grid[newX][newY].Signal = ele.Signal - 1
			}
			if grid[newX][newY].Visited == false && ele.Signal > 2 && grid[newX][newY].Signal != -1 {
				queue = append(queue, &grid[newX][newY])
			}
		}
	}

	fmt.Println(grid[x][y].Signal)
}
