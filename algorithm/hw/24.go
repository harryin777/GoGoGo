package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128990011
// 众数和中位数
func main24() {
	buf := bufio.NewScanner(os.Stdin)
	data := make([]int, 0, 100)
	for buf.Scan() {
		strs := strings.Split(buf.Text(), " ")
		for _, str := range strs {
			tmp, _ := strconv.Atoi(str)
			data = append(data, tmp)
		}
	}

	cal24(data)
}

func cal24(data []int) {
	dul := make(map[int]int)
	s := 0
	for _, datum := range data {
		dul[datum]++
		s = max(s, dul[datum])
	}

	newArr := make([]int, 0, len(data))
	for i, i2 := range dul {
		if i2 == s {
			newArr = append(newArr, i)
		}
	}
	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i] < newArr[j]
	})
	if len(newArr) == 1 {
		fmt.Println(newArr[0])
		return
	}
	if len(newArr)%2 == 0 {
		fmt.Println((newArr[len(newArr)/2] + newArr[(len(newArr)/2)-1]) / 2)
	} else {
		fmt.Println(newArr[len(newArr)/2])
	}
}
