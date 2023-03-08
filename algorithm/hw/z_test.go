package main

import (
	"fmt"
	"testing"
)

func TestCal5(t *testing.T) {
	//fmt.Println(cal8("3[m2[c]]"))
	//fmt.Println(cal8("3[k]2[mn]"))
	fmt.Println(cal8("1[k2[m3[n]]]"))
}

func TestCal9(t *testing.T) {
	cal9(3, [][]string{
		{"2", "5", "6", "7", "9", "5", "7"},
		{"1", "7", "4", "3", "4"},
	})
}

func TestCal10(t *testing.T) {
	//cal10([][]int{
	//	{0, 0, 0, -1, 0},
	//	{0, 0, 0, 0, 0},
	//	{0, 0, -1, 4, 0},
	//	{0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, -1},
	//	{0, 0, 0, 0, 0},
	//}, 1, 4)
}

func TestCal17(t *testing.T) {
	//fmt.Println(cal17("()(())"))
	cal17("()[(())]")
}

func TestCal18(t *testing.T) {
	cal18([]string{"1.5", "1.5.0"})
}

func TestPos(t *testing.T) {
	fmt.Println(pos(4))
}

func TestCal36(t *testing.T) {
	row, col := 5, 5
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
	opts := []int{1, 2, 2, 3, 3, 4, 4, 5}
	cal36(grid, opts, row, col)
}

func TestCal37(t *testing.T) {
	cal37([]int{2, 9, 11, 2, 3, 4, 13, 14, 7, 9, 14, 5, 6})
}

func TestBs(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(bs(0, len(data), 6, data))
}

func TestYIHUO(t *testing.T) {
	data := []int{6, 7}
	a1 := 5
	for i := 0; i < len(data); i++ {
		a1 ^= data[i]
		fmt.Println(a1)
	}
}
