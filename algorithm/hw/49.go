package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/129067003
// 高效的任务规划
func main49() {
	var count int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var machines int
		fmt.Scan(&machines)
		data := make([][]int, 0, machines)
		for j := 0; j < machines; j++ {
			var B, J int
			fmt.Scan(&B, &J)
			data = append(data, []int{B, J})
		}
		sort.Slice(data, func(i, j int) bool {
			if data[i][0] > data[j][0] {
				return true
			} else if data[i][0] == data[j][0] {
				return data[i][1] > data[j][0]
			} else {
				return false
			}
		})
		totalTime := 0
		remainTime := 0
		for p := 0; p < len(data); p++ {
			totalTime += data[p][0]
			remainTime -= data[p][0]
			if remainTime <= 0 {
				remainTime = data[p][1]
			} else {
				remainTime += data[p][0] - data[p][1]
			}

		}
		fmt.Println(totalTime + remainTime)
	}
}

func cal49() {

}
