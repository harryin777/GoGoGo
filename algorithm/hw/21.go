package main

import "fmt"

// https://dream.blog.csdn.net/article/details/128989969
// 字母消消乐
func main21() {
	var str string
	fmt.Scan(&str)
	cal21(str)
}

func cal21(str string) {
	left, right := 0, 1
	if len(str) == 2 {
		if str[0] == str[1] {
			fmt.Println(0)
		} else {
			fmt.Println(len(str))
		}
		return
	}

	for right < len(str) {
		if str[left] == str[right] {
			str = str[2:]
			left, right = 0, 1
		} else {
			left++
			right++
		}
	}
	fmt.Println(len(str))
}
