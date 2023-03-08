package main

import (
	"fmt"
)

// https://dream.blog.csdn.net/article/details/129031283
// 跳格子
func main42() {
	var count int
	fmt.Scan(&count)
	data := make([][]int, 0, 100)
	dul := make(map[int]int)
	endMap := make(map[int]struct{})
	startList := make([]int, 0, len(data))
	for {
		var one, two int
		n, _ := fmt.Scan(&one, &two)
		if n <= 0 {
			break
		}
		data = append(data, []int{one, two})
		startList = append(startList, one)
		dul[one]++
		dul[two]++
		endMap[two] = struct{}{}
	}

	// 找到起点，起点不能在被解锁格子的map里
	start := -1
	for i := 0; i < len(startList); i++ {
		if _, e := endMap[startList[i]]; !e {
			start = startList[i]
			break
		}
	}
	if start == -1 {
		fmt.Println("no")
		return
	}

	cal42(data, len(dul), start)
}

func cal42(data [][]int, show int, start int) {

	queue := make([]int, 0, len(data))
	queue = append(queue, start)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		show--
		// 找到一个以后，需要从list中剔除
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
