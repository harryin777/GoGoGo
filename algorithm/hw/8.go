package main

import (
	"fmt"
)

// https://dream.blog.csdn.net/article/details/128985488
// 压缩报文还原
func main8() {
	var str string
	fmt.Scan(&str)

	fmt.Println(cal8(str))
}

func cal8(str string) string {
	ans := ""
	tmp := ""
	stack := make([]string, 0, len(str))
	preNum := 1
	for i := 0; i < len(str); i++ {
		if str[i] <= '9' && str[i] >= '0' {
			preNum = int(str[i] - '0')
		}
		if str[i] == '[' {
			count := 1
			j := i + 1
			for count > 0 {
				if str[j] == '[' {
					count++
				} else if str[j] == ']' {
					count--
				}
				j++
			}
			tmp = cal8(str[i+1 : j-1])
			i = j - 1
		}

		if isChar(str[i]) {
			p := i + 1
			for p < len(str) && isChar(str[p]) {
				p++
			}
			tmp = str[i:p]
			i = p - 1
			for m := 0; m < preNum; m++ {
				stack = append(stack, tmp)
			}
			continue
		}

		if str[i] == ']' || len(str)-1 == i {
			for m := 0; m < preNum; m++ {
				stack = append(stack, tmp)
			}
		}
	}

	for i := 0; i < len(stack); i++ {
		ans += stack[i]
	}
	return ans
}

func isChar(b byte) bool {
	return b >= 'a' && b <= 'z'
}
