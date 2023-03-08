package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/129083500
// 寻找路径
func main59() {
	var arr []int
	mOne := int(^uint(0) >> 1)
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		arr = append(arr, 0)
		strs := strings.Split(buf.Text(), " ")
		for i := 0; i < len(strs); i++ {
			tmp, _ := strconv.Atoi(strs[i])
			if tmp != -1 && i != 0 {
				mOne = min(mOne, tmp)
			}
			arr = append(arr, tmp)
		}
	}

	cal59(arr, mOne)
}

func cal59(arr []int, mOne int) {
	ans := make([]int, 0, len(arr))
	var mPos int
	for i := 0; i < len(arr); i++ {
		if arr[i] == mOne && i != 0 {
			mPos = i
		}
	}
	for mPos >= 1 {
		ans = append(ans, arr[mPos])
		mPos /= 2
	}
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Printf("%d ", ans[i])
	}
}
