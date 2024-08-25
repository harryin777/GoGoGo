package main

import (
	"fmt"
)

/*
*
环中最长子串
*/
func main86() {
	var str string
	fmt.Scanln(&str)
	cal86(str)
}

func cal86(str string) {
	oCount := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'o' {
			oCount++
		}
	}

	if oCount%2 == 0 {
		fmt.Println(len(str))
	} else {
		fmt.Println(len(str) - 1)
	}
}
