package main

import (
	"fmt"
	"math"
)

// https://dream.blog.csdn.net/article/details/128994996
// 数字加减游戏
func main31() {
	var s, t, a, b int
	fmt.Scan(&s, &t, &a, &b)
	cal31(s, t, a, b)

}

func cal31(s, t, a, b int) {
	diff := math.Abs(float64(s) - float64(t))
	min1 := 0
	tmp := int(diff)
	for tmp%b != 0 {
		tmp -= a
		min1 += 1
	}
	min2 := 0
	tmp = int(diff)
	for tmp%b != 0 {
		tmp += a
		min2 += 1
	}
	fmt.Println(min(min2, min1))
}
