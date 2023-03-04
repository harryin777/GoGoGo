package main

import "fmt"

// https://dream.blog.csdn.net/article/details/128995115
// 新学校选址
func main32() {
	var houses int
	fmt.Scan(&houses)
	var hList []int
	maxOne := 0
	for i := 0; i < houses; i++ {
		var tmp int
		fmt.Scan(&tmp)
		maxOne = max(maxOne, tmp)
		hList = append(hList, tmp)
	}

	cal32(hList, maxOne)
}

func cal32(hList []int, maxOne int) {
	minP := 0
	total := int(^uint(0) >> 1)
	for i := 0; i <= maxOne; i++ {
		curr := 0
		for j := 0; j < len(hList); j++ {
			dist := hList[j] - i
			if dist < 0 {
				dist *= -1
			}
			curr += dist
		}
		if curr < total {
			total = curr
			minP = i
		}
	}
	fmt.Println(minP)
}
