package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/129018311
// 竖直四子棋
func main36() {
	var row, col int
	fmt.Scan(&row, &col)
	buf := bufio.NewScanner(os.Stdin)
	opts := make([]int, 0, col*row)
	for buf.Scan() {
		data := buf.Text()
		dataArr := strings.Split(data, " ")
		for i := 0; i < len(dataArr); i++ {
			if len(dataArr[i]) == 0 {
				continue
			}
			tmp, _ := strconv.Atoi(dataArr[i])
			opts = append(opts, tmp)
		}
	}
	grid := make([][]QI, row)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]QI, col)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = QI{
				X: i,
				Y: j,
			}
		}
	}

	cal36(grid, opts, row, col)
}
func cal36(grid [][]QI, opts []int, row, column int) {
	actions := [][]int{{1, 0}, {0, 1}, {1, 1}, {-1, -1}}
	for i := 0; i < len(opts); i++ {
		color := 0
		if i%2 == 0 {
			color = 1
		} else {
			color = 2
		}
		if opts[i]-1 > len(grid) || opts[i]-1 < 0 {
			fmt.Printf("%d,error", i+1)
			return
		}
		if len(grid[opts[i]-1]) > column {
			fmt.Printf("%d,error", i+1)
			return
		}
		for p := 0; p < len(grid[opts[i]-1]); p++ {
			if grid[opts[i]-1][p].Color == 0 {
				grid[opts[i]-1][p].Color = color
				break
			}
		}

		for m := 0; m < row; m++ {
			for n := 0; n < column; n++ {
				for a := 0; a < len(actions); a++ {
					count := 1
					preColor := 0
					x, y := m, n
					for x >= 0 && x < row && y >= 0 && y < column {
						if grid[x][y].Color == 0 {
							goto next
						}
						if preColor == 0 {
							preColor = grid[x][y].Color
						} else {
							if preColor != grid[x][y].Color {
								count = 0
								break
							} else {
								count++
							}
						}
						x += actions[a][0]
						y += actions[a][1]
						if count == 4 {
							if preColor == 1 {
								fmt.Printf("%d,red", i+1)
							} else {
								fmt.Printf("%d,blue", i+1)
							}
							return
						}
					}

				}
			next:
			}
		}
	}

}

type QI struct {
	X     int
	Y     int
	Color int // 1 红，2蓝
}
