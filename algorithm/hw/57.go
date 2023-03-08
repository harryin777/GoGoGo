package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129052829#comments_25443921
// 勾股数
func main57() {
	var n, m int
	fmt.Scan(&n, &m)
	cal57(n, m)
}

func cal57(n, m int) {
	var count int
	for i := n; i <= m; i++ {
		for j := i + 1; j <= m; j++ {
			for k := j + 1; k <= m; k++ {
				if relatively_prime(i, j) && relatively_prime(i, k) && relatively_prime(j, k) && (i*i+j*j == k*k) {
					fmt.Printf("%d %d %d \n", i, j, k)
					count++
				}
			}
		}
	}

	if count == 0 {
		fmt.Println("Na")
	}
}

// 是否互质的求法
func relatively_prime(x, y int) bool {
	if x == y && y == 1 {
		return false
	}
	mOne := min(x, y)
	for i := 2; i <= mOne; i++ {
		if x%i == 0 && y%i == 0 {
			return false
		}
	}

	return true
}
