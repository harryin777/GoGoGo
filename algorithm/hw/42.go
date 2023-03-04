package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/129031283
// 跳格子
func main42() {
	var count int
	fmt.Scan(&count)
	data := make([][]int, 0, 100)
	dul := make(map[int]int)
	for {
		var one, two int
		n, _ := fmt.Scan(&one, &two)
		if n <= 0 {
			break
		}
		data = append(data, []int{one, two})
		dul[one]++
		dul[two]++
	}

	cal42(data, len(dul))
}

func cal42(data [][]int, show int) {
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	for i := 0; i < len(data); i++ {
		if data[0][0] == data[i][1] {
			fmt.Println("no")
			return
		}
	}

	queue := make([]int, 0, len(data))
	queue = append(queue, data[0][0])
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		show--
		var tmp [][]int
		for i := 0; i < len(data); i++ {
			if data[i][0] == cur {
				queue = append(queue, data[i][1])
			} else {
				tmp = append(tmp, data[i])
			}
		}
		data = tmp
	}
	if show == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
