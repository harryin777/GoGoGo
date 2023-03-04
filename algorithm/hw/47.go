package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129095331
// 九宫格按键输入
func main47() {
	var str string
	fmt.Scan(&str)
	cal47(str)
}

func cal47(command string) {
	numMapC := map[int]string{
		1: ",.",
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	ans := ""
	curType := 1
	for i := 0; i < len(command); i++ {
		if command[i] == '#' {
			curType ^= 1
			continue
		}
		if command[i] == '/' {
			continue
		}
		if curType == 1 {
			ans += string(command[i])
		} else {
			preNum := command[i] - '0'
			j := i + 1
			for j < len(command) {
				if command[j] == command[i] {
					j++
				} else {
					break
				}
			}
			pos := j - i - 1
			if pos >= len(numMapC[int(preNum)]) {
				pos -= len(numMapC[int(preNum)])
			}
			ans += string(numMapC[int(preNum)][pos])
			i = j - 1
		}
	}
	fmt.Println(ans)
}
