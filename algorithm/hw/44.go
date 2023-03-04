package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129103070
// 旋转骰子
func main44() {
	var command string
	fmt.Scan(&command)
	cal44(command)
}

func cal44(command string) {
	l, r, f, b, u, d := 1, 2, 3, 4, 5, 6
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case 'L':
			l, r, u, d = u, d, r, l
		case 'R':
			l, r, u, d = d, u, l, r
		case 'F':
			f, b, u, d = u, d, b, f
		case 'B':
			f, b, u, d = d, u, f, b
		case 'C':
			l, r, f, b = f, b, r, l
		case 'A':
			l, r, f, b = b, f, l, r
		}
	}
	fmt.Printf("%v%v%v%v%v%v", l, r, f, b, u, d)
}
