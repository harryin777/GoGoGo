package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
分配土地最大面积
*/

func main77() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	buf := bufio.NewScanner(os.Stdin)
	matrix := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		buf.Scan()
		arr := strings.Split(buf.Text(), " ")
		arrS := make([]int, 0, len(arr))
		for j := 0; j < len(arr); j++ {
			val, _ := strconv.Atoi(arr[j])
			arrS = append(arrS, val)
		}
		matrix = append(matrix, arrS)
	}
	cal77(matrix)
}

func cal77(matrix [][]int) {
	flagMapCordinateSlice := make(map[int][][]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				continue
			}
			flagMapCordinateSlice[matrix[i][j]] = append(flagMapCordinateSlice[matrix[i][j]], []int{i, j})
		}
	}

	ans := 0
	for _, val := range flagMapCordinateSlice {
		left, right, top, bottom := 600, -1, 600, -1
		for _, cordinates := range val {
			if cordinates[0] < top {
				top = cordinates[0]
			}
			if cordinates[0] > bottom {
				bottom = cordinates[0]
			}
			if cordinates[1] < left {
				left = cordinates[1]
			}
			if cordinates[1] > right {
				right = cordinates[1]
			}
			area := (right - left + 1) * (bottom - top + 1)
			if area > ans {
				ans = area
			}
		}
	}

	fmt.Println(ans)
}
