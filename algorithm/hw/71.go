package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*
题目描述与示例
题目描述
已知火星人使用的运算符号为#、
他们与地球人的等价公式如下：
．x#y=4*x+3*y+2
其中xY是无符号整数Q
地球人公式按照c语言规则进行计算
火星人公式中#符优先级高于$
相同的运算符按从左到右的顺序运算

输入描述
火星人字符串表达式结尾不带回车换行
*/
func main71() {
	buf := bufio.NewScanner(os.Stdin)
	if buf.Scan() {
		str := buf.Text()
		cal71(str)
	}
}

// 7#6$5#12
func cal71(str string) {
	num := 0
	stack := make([]int, 0, 10)
	presign := "$"
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			num = num*10 + int(str[i]-'0')

		}
		if i == len(str)-1 || str[i] == '#' || str[i] == '$' {
			if presign == "#" {
				stack[len(stack)-1] = stack[len(stack)-1]*4 + 3*num + 2
			} else if presign == "$" {
				stack = append(stack, num)
			}
			presign = string(str[i])
			num = 0
		}
	}

	for len(stack) > 1 {
		tmp := stack[0]*2 + stack[1] + 3
		stack = stack[2:]
		stack = append([]int{tmp}, stack...)
	}

	fmt.Println(stack[0])
}
