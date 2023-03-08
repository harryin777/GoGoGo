package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129052681
// 最大报酬
func main61() {
	var minutes, count int
	tasks := make([]task61, 0, count)
	fmt.Scan(&minutes, &count)
	for i := 0; i < count; i++ {
		var time, salary int
		fmt.Scan(&time, &salary)
		tasks = append(tasks, task61{
			Time:   time,
			Salary: salary,
		})
	}

	cal61(tasks, minutes)
}

func cal61(tasks []task61, minutes int) {
	dp := make([][]int, len(tasks)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, minutes+1)
	}

	for i := 1; i <= len(tasks); i++ {
		for j := 1; j <= minutes; j++ {
			if j < tasks[i-1].Time {
				dp[i][j] = dp[i-1][j]
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-tasks[i-1].Time]+tasks[i-1].Salary)
		}
	}

	fmt.Println(dp[len(tasks)][minutes])
}

type task61 struct {
	Time   int
	Salary int
}
