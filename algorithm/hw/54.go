package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129004798
// 猴子爬山
func main54() {
	var n int
	fmt.Scan(&n)
	cal54(n)
}

func cal54(n int) {
	dp1 := 1
	dp2 := 1
	dp3 := 2
	ans := 0
	if n == 0 {
		///
	}
	for i := 3; i < n; i++ {
		ans = dp3 + dp1
		dp1 = dp2
		dp2 = dp3
		dp3 = ans
	}
	fmt.Println(ans)
}
