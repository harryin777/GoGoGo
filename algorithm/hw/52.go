package main

import (
	"fmt"
	"strconv"
)

// https://dream.blog.csdn.net/article/details/128985324
// 火星人操作符
func main52() {
	var com string
	fmt.Scan(&com)
	cal52(com)
}

func cal52(com string) {
	stack := make([]int, 0, len(com))
	preSign := 0
	ans := 0

	for i := 0; i < len(com); i++ {
		if isDigit(com[i]) {
			count := i + 1
			for count < len(com) {
				if isDigit(com[count]) {
					count++
				} else {
					break
				}
			}
			tmpInt, _ := strconv.Atoi(com[i:count])
			if preSign == 0 {
				stack = append(stack, tmpInt)
			} else {
				stack[len(stack)-1] = sharp(stack[len(stack)-1], tmpInt)
				preSign = 0
			}

			i = count - 1
		}
		if com[i] == '#' {
			preSign = 1
		}
	}

	for i := 0; i < len(stack)-1; i++ {
		tmp := dollar(stack[i], stack[i+1])
		ans = tmp
		stack[i+1] = tmp
	}

	fmt.Println(ans)
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

func sharp(x, y int) int {
	return 4*x + 3*y + 2
}

func dollar(x, y int) int {
	return 2*x + y + 3
}
