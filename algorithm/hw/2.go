package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/129075162
// 流水线
func main2() {
	var m, taskCount int
	fmt.Scan(&m, &taskCount)
	taskList := make([]int, 0, taskCount)
	for i := 0; i < taskCount; i++ {
		var tmp int
		fmt.Scan(&tmp)
		taskList = append(taskList, tmp)
	}
	cal1(m, taskList)
}

func cal1(m int, taskList []int) {
	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i] < taskList[j]
	})
	table := make([][]int, m)
	for i := 0; i < len(table); i++ {
		table[i] = make([]int, 1, 1)
	}

	for i := 0; i < len(table); i++ {
		if table[i][0] == 0 && len(taskList) != 0 {
			table[i][0] = taskList[0]
			taskList = taskList[1:]
		}
	}

	mIndex := 0
	for len(taskList) != 0 {
		task := taskList[0]
		taskList = taskList[1:]
		table[mIndex][0] += task
		mIndex++

	}

	ans := 0
	for i := 0; i < len(table); i++ {
		ans = max(ans, table[i][0])
	}

	fmt.Println(ans)
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
