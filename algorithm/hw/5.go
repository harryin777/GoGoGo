package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main5() {
	var a int
	fmt.Scan(&a)
	seats := make([]int, 0, a)
	var str string
	fmt.Scanf("[%v]", str)
	strs := strings.Split(str, ",")
	for i := 0; i < len(strs); i++ {
		seatInt, _ := strconv.Atoi(strs[i])
		seats = append(seats, seatInt)
	}
	cal5(seats, a)
}

func cal5(people []int, seatCount int) {
	seat := make([]int, seatCount)
	ans := 0
	for i := 0; i < len(people); i++ {
		if i == 0 {
			seat[0] = 1
			continue
		}
		if i == 1 {
			seat[len(seat)-1] = 1
		}
		if people[i] < 0 {
			seat[-(people[i])] = 0
			continue
		}

		pos := 0
		maxLenIndex, maxLen := 0, 0
		for pos < len(seat) {
			if seat[pos] != 0 {
				pos++
				continue
			}
			left, right := pos, pos
			leftLength, rightLength := 0, 0
			for left >= 0 && seat[left] == 0 {
				leftLength++
				left--
			}
			for right < len(seat) && seat[right] == 0 {
				rightLength++
				right++
			}

			mm := min(leftLength, rightLength)
			if mm > maxLen {
				maxLen = mm
				maxLenIndex = pos
				ans = maxLenIndex
			}
			pos++
		}
		seat[maxLenIndex] = 1
	}

	fmt.Println(ans + 1)
}
