package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129031325
// 分糖果
func main56() {
	var count int
	fmt.Scan(&count)
	cal56(count)
}

func cal56(count int) {
	var ans int
	for count != 1 {
		if count == 3 {
			ans += 2
			break
		}
		if count%2 != 0 {
			if (count+1)/2%2 == 0 {
				count += 1
			} else {
				count -= 1
			}
			ans += 1
		}
		count /= 2
		ans += 1
	}

	fmt.Println(ans)
}
