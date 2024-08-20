package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
*反射计数
 */
func main78() {
	var w, h, x, y, sx, sy, t int
	fmt.Scanf("%d %d %d %d %d %d %d", &w, &h, &x, &y, &sx, &sy, &t)
	buf := bufio.NewScanner(os.Stdin)
	matrix := make([][]string, 0, h)
	for i := 0; i < h; i++ {
		buf.Scan()
		matrix = append(matrix, strings.Split(buf.Text(), ""))
	}
	cal78(matrix, y, x, sy, sx, t)
}

func cal78(matrix [][]string, x, y, sx, sy, t int) {
	ans := 0
	if matrix[x][y] == "1" {
		ans++
	}
	curX, curY := x, y
	for i := 0; i < t; i++ {
		nx, ny := curX+sx, curY+sy
		if nx < 0 || nx > len(matrix)-1 {
			sx = -sx
			nx = curX + sx
		}
		if ny < 0 || ny > len(matrix[0])-1 {
			sy = -sy
			ny = curY + sy
		}
		if matrix[nx][ny] == "1" {
			ans++
		}
		curX, curY = nx, ny
	}

	fmt.Println(ans)
}
