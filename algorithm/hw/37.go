package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/129019175
// 斗地主2
func main37() {
	buf := bufio.NewScanner(os.Stdin)
	pokers := make([]int, 0, 13)
	for buf.Scan() {
		data := buf.Text()
		arr := strings.Split(data, " ")
		for i := 0; i < len(arr); i++ {
			if arr[i] == "J" {
				pokers = append(pokers, 11)
				continue
			} else if arr[i] == "Q" {
				pokers = append(pokers, 12)
				continue
			} else if arr[i] == "K" {
				pokers = append(pokers, 13)
				continue
			} else if arr[i] == "A" {
				pokers = append(pokers, 14)
				continue
			}
			tmp, _ := strconv.Atoi(arr[i])
			pokers = append(pokers, tmp)
		}
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
		pre := i
		var j = i + 1
		for j = i + 1; j < len(pokers); j++ {
			if pokers[j]-pokers[pre] == 1 {
				pre = j
			} else {
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
		fmt.Println()
	}
}
