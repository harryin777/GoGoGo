package main

import (
	"fmt"
	"strconv"
)

/*
*
螺旋数字矩阵
*/
func main74() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	cal74(n, m)
}

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func cal74(n, m int) {
	column := n/m + 1
	matrix := make([][]string, m)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]string, column)
	}
	i, j, total := 0, 0, m*column
	for k := 1; k <= total; k++ {
		count, curDir := 0, 0
		x, y := i, j
		for i >= m || i < 0 || j >= column || j < 0 || matrix[i][j] != "" {
			if count == 4 {
				goto here
			}
			i, j = x+dir[curDir%4][0], y+dir[curDir%4][1]
			curDir++
			count++
		}
		if k <= n {
			matrix[i][j] = strconv.Itoa(k)
		} else {
			matrix[i][j] = "*"
		}
	}
here:
	for _, col := range matrix {
		for _, val := range col {
			fmt.Printf("%v ", val)
		}
		fmt.Println()
	}
}
