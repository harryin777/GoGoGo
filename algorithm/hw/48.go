package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/129066972
// 分积木
func main48() {
	var count int
	fmt.Scan(&count)
	sum := 0
	data := make([]int, 0, count)
	for i := 0; i < count; i++ {
		var t int
		fmt.Scan(&t)
		sum += t
		data = append(data, t)
	}
	cal48(data, sum)
}

func cal48(data []int, total int) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	mmOne := data[0]
	xor := mmOne
	for i := 1; i < len(data); i++ {
		xor ^= data[i]
	}
	if xor != 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println(total - mmOne)
}
