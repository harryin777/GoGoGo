package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128985996
// 斗地主
func main30() {
	var inS, outS string
	fmt.Scan(&inS, &outS)
	var in, out []int
	tmp := strings.Split(inS, "-")
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == "J" {
			in = append(in, 11)
			continue
		} else if tmp[i] == "Q" {
			in = append(in, 12)
			continue
		} else if tmp[i] == "K" {
			in = append(in, 13)
			continue
		} else if tmp[i] == "A" {
			in = append(in, 14)
			continue
		}
		tt, _ := strconv.Atoi(tmp[i])
		in = append(in, tt)
	}
	tmp1 := strings.Split(outS, "-")
	for i := 0; i < len(tmp1); i++ {
		if tmp1[i] == "J" {
			out = append(out, 11)
			continue
		} else if tmp1[i] == "Q" {
			out = append(out, 12)
			continue
		} else if tmp1[i] == "K" {
			out = append(out, 13)
			continue
		} else if tmp1[i] == "A" {
			out = append(out, 14)
			continue
		}
		tt, _ := strconv.Atoi(tmp1[i])
		out = append(out, tt)
	}
	cal30(in, out)
}

func cal30(in, out []int) {
	grid := make([]int, 12)
	for i := 0; i < len(grid); i++ {
		grid[i] = 4
	}
	// 减三是因为最小牌是3，对应索引是0
	for i := 0; i < len(in); i++ {
		grid[in[i]-3]--
	}
	for i := 0; i < len(out); i++ {
		grid[out[i]-3]--
	}
	mlen := 5
	finalLeft, finalRight := 0, 0
	for i := 0; i < len(grid); i++ {
		left, right := 0, 0
		for j := i; j < len(grid); j++ {
			if j == i {
				left = j
			}
			if grid[j] == 0 {
				left = j + 1
			}
			if j-i >= 5 {
				right = j
			}
		}
		if right-left >= mlen {
			mlen = right - left
			finalLeft, finalRight = left, right
		}
	}
	ans := ""
	if finalRight-finalLeft < 5 {
		fmt.Println("NO-CHAIN")
		return
	}
	for i := finalLeft; i <= finalRight; i++ {
		if i == 8 {
			ans += "J-"
			continue
		} else if i == 9 {
			ans += "Q-"
			continue
		} else if i == 10 {
			ans += "K-"
			continue
		} else if i == 11 {
			ans += "A-"
			continue
		}
		ans += strconv.FormatInt(int64(i+3), 10) + "-"
	}
	fmt.Println(ans[:len(ans)-1])
}
