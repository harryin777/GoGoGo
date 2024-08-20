package main

import "fmt"

/*
*
“吃货"和“馋嘴"两人到披萨店点了一份铁盘（圆形）披萨，并嘱咐店员将披萨按放射状切成大小相同的偶数扇形小块。但
是粗心服务员将技萨切成了每块大小都完全不同〔奇数块〗，且肉眼能分辨出大小。
由于两人都想吃到最多的披萨，他们商量了一个他们认为公平的分法：从“吃货"开始，轮流取披萨。除了第一块披萨可以
任意选取以夕卜，其他都必须从缺口开始选。他俩选披萨的思路不同。
“馋嘴"每次都会选最大块的披萨，而且“吃货"知道“馋嘴"的相法

已知披萨小块的数量以及每块的大小，求“吃货"能分得的最大的披萨大小的总和。
输入描述
第1行为一个正整数奇数，表示披萨小块数量。3<=N<=5。
接下来的第2行到第N+I行（共行),每行为一个正整数，表示第I块披萨的大小。1<=i<=N。
披萨小块从某一块开始，按照一个方向依次顺序编号为。每块披萨的大小范围为[1, 21474836471]。
输出描述
“吃货"能分得的最大的披萨大小的总和。

5
8
2
10
5
7
*/
func main65() {
	var pieces int
	fmt.Scanln(&pieces)
	var arr []int
	for i := 0; i < pieces; i++ {
		var val int
		fmt.Scan(&val)
		arr = append(arr, val)
	}
	cal65(arr)
}

func cal65(arr []int) {
	maxVal := 0
	k := 0
	for index, i := range arr {
		if i > maxVal {
			maxVal = i
			k = index
		}
	}
	nArr := make([]int, 0, len(arr))
	for i := 0; i < k; i++ {
		nArr = append(nArr, arr[i])
	}

	nArr = append(arr[k+1:], nArr...)
	for len(nArr) != 0 {
		if nArr[0] > nArr[len(nArr)-1] {
			maxVal += nArr[len(nArr)-1]
		} else if nArr[0] > nArr[len(nArr)-1] {
			break
		} else {
			maxVal += nArr[0]
		}
		nArr = nArr[1:]
		nArr = nArr[:len(nArr)-1]
	}

	fmt.Println(maxVal)
}
