package main

import "fmt"

/**
变换最小字符串
*/

func main() {
	var str string
	fmt.Scanln(&str)
	cal88(str)
}

func cal88(str string) {
	lastOne := -1
	for i := len(str) - 1; i > 0; i++ {
		if str[i] < str[i-1] {
			lastOne = i
		}
	}

	if lastOne == -1 {
		fmt.Println(str)
	}
	preOne := -1
	for i := 0; i < len(str)-2; i++ {
		if str[i] > str[i+1] {
			preOne = i
		}
	}
	tmp := []byte(str)
	tmp[preOne], tmp[lastOne] = tmp[lastOne], tmp[preOne]
	fmt.Println(string(tmp))
}
