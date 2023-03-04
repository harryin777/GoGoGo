package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128986142
// 出租车计费
func main51() {
	var targetNum int
	fmt.Scan(&targetNum)
	cal51(targetNum)
	//cal51O(targetNum)
}

func cal51(t int) {
	ans := 0

	for i := 1; i <= t; i++ {
		str := strconv.FormatInt(int64(i), 10)
		if strings.Contains(str, "4") {
			continue
		}
		ans++
	}

	fmt.Println(ans)
}

func cal51O(t int) {
	ans, tmp, k, j := t, 0, 0, 1

	for t > 0 {
		if t%4 > 0 {
			tmp += (t%10-1)*k + j
		} else {
			tmp += (t % 10) * k
		}
		k = k*9 + j
		j *= 10
		t /= 10
	}

	fmt.Println(ans - tmp)
}
