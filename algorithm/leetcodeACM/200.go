package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	grid := make([][]string, 0, 10)
	for buf.Scan() {
		arr := strings.Split(buf.Text(), " ")
		grid = append(grid, arr)
	}

	fmt.Println(islandCount(grid))
}

var dirArr = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func islandCount(grid [][]string) int {
	ans := 0
	visitedMap := make([][]bool, len(grid))
	for i := 0; i < len(visitedMap); i++ {
		visitedMap[i] = make([]bool, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "1" {
				ans++
				queue := [][]int{
					{i, j},
				}
				visitedMap[i][j] = true
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					for k := 0; k < len(dirArr); k++ {
						nx, ny := cur[0]+dirArr[k][0], cur[1]+dirArr[k][1]
						if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[i]) {
							continue
						}
						if visitedMap[nx][ny] {
							continue
						}
						if grid[nx][ny] == "1" {
							queue = append(queue, []int{nx, ny})
							visitedMap[nx][ny] = true
							grid[nx][ny] = "0"
						}
					}
				}
			}
		}
	}

	return ans
}
