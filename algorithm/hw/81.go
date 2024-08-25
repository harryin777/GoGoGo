package main

import (
	"fmt"
)

/**
密码解密
*/

func main81() {
	var str string
	fmt.Scanln(&str)
	cal81(str)
}

func cal81(str string) {
	ans := ""
	for i := 0; i < len(str); i++ {
		if str[i] != '*' {
			ans += string(str[i] - '0' + 'a' - 1)
		} else if i > 1 {
			ans = ans[:len(ans)-2]
			ans += string((str[i-2]-'0')*10 + str[i-1] - '0' + 'a' - 1)
		}
	}

	fmt.Println(ans)

}
