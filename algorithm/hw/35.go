package main

import (
	"fmt"
	"strings"
)

// https://dream.blog.csdn.net/article/details/129004861
// 找单词
func main35() {
	var r int
	fmt.Scan(&r)
	grid := make([][]Ch, 0, r)
	for i := 0; i < r; i++ {
		var str string
		fmt.Scan(&str)
		strs := strings.Split(str, ",")
		row := make([]Ch, 0, len(strs))
		for j := 0; j < len(strs); j++ {
			row = append(row, Ch{
				X:       i,
				Y:       j,
				CH:      strs[j],
				Visited: false,
			})
		}
		grid = append(grid, row)
	}
	var word string
	fmt.Scan(&word)
	cal35(grid, word)
}

func cal35(grid [][]Ch, word string) {
	ans := make([][]int, 0, len(word))
	var dfs func(int, int, int) bool
	dfs = func(currIndex int, x int, y int) bool {
		if currIndex >= len(word) || x < 0 || x >= len(grid) || y < 0 || y > len(grid) || grid[x][y].Visited {
			return false
		}
		grid[x][y].Visited = true
		if string(word[currIndex]) != grid[x][y].CH {
			return false
		}
		if currIndex == len(word)-1 {
			ans = append(ans, []int{x, y})
			return true
		}
		if dfs(currIndex+1, x+1, y) {
			ans = append(ans, []int{x, y})
			return true
		}
		if dfs(currIndex+1, x-1, y) {
			ans = append(ans, []int{x, y})
			return true
		}
		if dfs(currIndex+1, x, y+1) {
			ans = append(ans, []int{x, y})
			return true
		}
		if dfs(currIndex+1, x, y-1) {
			ans = append(ans, []int{x, y})
			return true
		}
		return false
	}

	dfs(0, 0, 0)
	if len(ans) == 0 {
		fmt.Println(word)
	} else {
		ansStr := ""
		for i := len(ans) - 1; i >= 0; i-- {
			ansStr = ansStr + fmt.Sprintf("%d,%d,", ans[i][0], ans[i][1])
		}
		fmt.Println(ansStr[:len(ansStr)-1])
	}
}

type Ch struct {
	X       int
	Y       int
	CH      string
	Visited bool
}

func cal35t1(grid [][]string, word string) {
	ans := make([][]int, 0, len(word))
	var dfs func(int, int, int, int) bool
	// directFrom 1从左往右， 2是从右往左， 3是从上往下， 4是从下往上
	// 这种写法避免了找回到头路，但是代码量有点大
	dfs = func(currIndex int, x int, y int, directFrom int) bool {
		if currIndex >= len(word) || x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return false
		}
		if string(word[currIndex]) != grid[x][y] {
			return false
		}

		if currIndex == len(word)-1 {
			ans = append(ans, []int{x, y})
			return true
		}

		if directFrom == 1 {
			if dfs(currIndex+1, x, y+1, 1) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x+1, y, 3) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x-1, y, 4) {
				ans = append(ans, []int{x, y})
				return true
			}
		} else if directFrom == 2 {
			if dfs(currIndex+1, x, y-1, 2) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x+1, y, 3) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x-1, y, 4) {
				ans = append(ans, []int{x, y})
				return true
			}
		} else if directFrom == 3 {
			if dfs(currIndex+1, x, y+1, 1) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x+1, y, 3) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x, y-1, 2) {
				ans = append(ans, []int{x, y})
				return true
			}
		} else if directFrom == 4 {
			if dfs(currIndex+1, x, y+1, 1) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x, y-1, 2) {
				ans = append(ans, []int{x, y})
				return true
			}
			if dfs(currIndex+1, x-1, y-1, 4) {
				ans = append(ans, []int{x, y})
				return true
			}
		}
		return false
	}
	dfs(0, 0, 0, 1)
	if len(ans) == 0 {
		fmt.Println(word)
	} else {
		ansStr := ""
		for i := len(ans) - 1; i >= 0; i-- {
			ansStr = ansStr + fmt.Sprintf("%d,%d,", ans[i][0], ans[i][1])
		}
		fmt.Println(ansStr[:len(ansStr)-1])
	}
}
