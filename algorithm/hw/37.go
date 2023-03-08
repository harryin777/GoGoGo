package main

import (
	"fmt"
	"sort"
	"strconv"
)

// https://dream.blog.csdn.net/article/details/129019175
// 斗地主2
func main37() {
	pokers := make([]int, 0, 13)
	for {
		var str string
		n, _ := fmt.Scan(&str)
		if n <= 0 {
			break
		}
		if str == "J" {
			pokers = append(pokers, 11)
			continue
		} else if str == "Q" {
			pokers = append(pokers, 12)
			continue
		} else if str == "K" {
			pokers = append(pokers, 13)
			continue
		} else if str == "A" {
			pokers = append(pokers, 14)
			continue
		}
		tmp, _ := strconv.Atoi(str)
		pokers = append(pokers, tmp)
	}
	cal37(pokers)
}

func cal37(pokers []int) {
	sort.Slice(pokers, func(i, j int) bool {
		return pokers[i] < pokers[j]
	})
	ans := make([][]int, 0, 10)
	for i := 0; i < len(pokers); i++ {
		if pokers[i] == 2 {
			continue
		}
		var j = i + 1
		for j = i + 1; j < len(pokers); j++ {
			if pokers[j]-pokers[j-1] != 1 {
				break
			}
		}
		if j-i >= 5 {
			ans = append(ans, pokers[i:j])

		}
		i = j - 1
	}
	if len(ans) == 0 {
		fmt.Println("NO")
		return
	}

	for i := 0; i < len(ans); i++ {
		//fmt.Println(ans[i])
		for _, val := range ans[i] {
			fmt.Printf("%v ", val)
		}
		fmt.Printf("\n")
	}
}
