package main

import (
	"fmt"
	"strings"
)

/*
*
题目描述
给一个正整数
输入
输出
输出一个数字字符串，记录最小值。
NUMI，
计算出新正整数NlJM2。
NlJM2为NIJMI中移除N位数字后的结果，需要使得的值最小。
1，输入的第一行为一个字符串，字符串由9一9字符组成，记录正整数，
2，输入的第二行为需要移除的数字的个数，小于NIJMI长度。
长度小于32。
*/
func main69() {
	var num string
	fmt.Scanln(&num)
	var k int
	fmt.Scan(&k)
	cal69(num, k)
}

func cal69(num string, k int) {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}

	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		fmt.Println("0")
	}

	fmt.Println(ans)
}
