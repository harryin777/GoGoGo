package main

import "fmt"

/**
提取字符串中的最长合法简单数学表达式，字符串长度最长的，并计算表达式的值。如果没有，则返回。
简单数学表达式只能包含以下内容：9数字，符号+一*
说明：
1，所有数字，计算结果都不超过long
2，如果有多个长度一样的，请返回第一个表达式的结果
3，数学表达式，必须是最长的，合法的
4，操作符不能连续出现，如+一+1是不合法的
输入描述
字符串
输出描述
表达式值
*/

func main68() {
	var str string
	fmt.Scanln(&str)
	cal68(str)
}

func calculator(str string) int {
	preSign := "+"
	stack := make([]int, 0, 10)
	num, ans := 0, 0
	for i := 0; i < len(str); i++ {
		var isDigit bool
		if str[i] >= '0' && str[i] <= '9' {
			isDigit = true
			num = int(str[i]-'0') + num*10
		}
		if !isDigit || i == len(str)-1 {
			switch preSign {
			case "+":
				stack = append(stack, num)
			case "-":
				stack = append(stack, -num)
			case "*":
				stack[len(stack)-1] *= num
			case "/":
				stack[len(stack)-1] /= num
			}
			num = 0
			preSign = string(str[i])
		}
	}

	for _, i := range stack {
		ans += i
	}

	return ans
}

// 1+2+3
func cal68(str string) {
	wrongIndex := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		if !(str[i] == '-' || str[i] == '+' || str[i] == '*') && !(str[i] >= '0' && str[i] <= '9') {
			wrongIndex[i] = 1
		}
	}
	curStr := ""
	for i := 0; i < len(str); i++ {
		if wrongIndex[i] == 1 {
			continue
		}
		for j := i + 1; j <= len(str); j++ {
			tmp := str[i:j]
			if check(tmp) {
				if j-i > len(curStr) {
					curStr = str[i:j]
				}
			}
		}
	}
	fmt.Println(calculator(curStr))
}

func check(str string) bool {
	preSign := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		var isSymbol, isDigital bool
		if str[i] == '-' || str[i] == '+' || str[i] == '*' {
			isSymbol = true
			preSign[i] = str[i]
		}
		if str[i] >= '0' && str[i] <= '9' {
			isDigital = true
		}
		if !isDigital && !isSymbol {
			return false
		}
	}
	for i := 0; i < len(preSign)-1; i = i + 2 {
		if preSign[i] != 0 && preSign[i+1] != 0 {
			return false
		}
	}
	if preSign[len(preSign)-1] != 0 {
		return false
	}

	return true
}
