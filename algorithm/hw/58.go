package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129045100
// 开心消消乐
func main58() {
	var r, c int
	fmt.Scan(&r, &c)
	grid := make([][]int, 0, r)
	for i := 0; i < r; i++ {
		data := make([]int, 0, c)
		for j := 0; j < c; j++ {
			var tmp int
			fmt.Scan(&tmp)
			data = append(data, tmp)
		}
		grid = append(grid, data)
	}

	cal58(grid)
}

func cal58(grid [][]int) {
	ans := 0
	actions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				ans++
				queue := make([][]int, 0, len(grid))
				queue = append(queue, []int{i, j})
				for len(queue) != 0 {
					cur := queue[0]
					queue = queue[1:]
					for a := 0; a < len(actions); a++ {
						xNew, yNew := cur[0]+actions[a][0], cur[1]+actions[a][1]
						if xNew < 0 || xNew >= len(grid) || yNew < 0 || yNew >= len(grid) {
							continue
						}
						if grid[xNew][yNew] == 1 {
							queue = append(queue, []int{xNew, yNew})
						}
						grid[xNew][yNew] = 0
					}
				}
			}
		}
	}

	fmt.Println(ans)
}
