package main

import "fmt"

// https://dream.blog.csdn.net/article/details/128989874
// 括号检查
func main17() {
	var a string
	fmt.Scan(&a)

	cal17(a)
}

func cal17(str string) {
	maxDepth := 0
	currDepth := 0
	queue := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			currDepth++
			queue = append([]byte{'('}, queue...)
		} else if str[i] == '{' {
			currDepth++
			queue = append([]byte{'{'}, queue...)
		} else if str[i] == '[' {
			currDepth++
			queue = append([]byte{'['}, queue...)
		}

		if str[i] == ')' {
			if len(queue) == 0 || queue[0] != '(' {
				fmt.Println(0)
				return
			} else if queue[0] == '(' {
				queue = queue[1:]
				currDepth--
			}
		}
		if str[i] == ']' {
			if len(queue) == 0 || queue[0] != '[' {
				fmt.Println(0)
				return
			} else if queue[0] == '[' {
				queue = queue[1:]
				currDepth--
			}
		}
		if str[i] == '}' {
			if len(queue) == 0 || queue[0] != '{' {
				fmt.Println(0)
				return
			} else if queue[0] == '{' {
				queue = queue[1:]
				currDepth--
			}
		}

		maxDepth = max(maxDepth, currDepth)
	}
	fmt.Println(maxDepth)
}

//func cal17(str string) int {
//	maxDepth := 0
//	currDepth := 0
//	for i := 0; i < len(str); i++ {
//		if str[i] == '(' {
//			currDepth++
//			count := 1
//			j := i + 1
//			for count > 0 {
//				if j >= len(str) {
//					return 0
//				}
//				if str[j] == '(' {
//					count++
//				} else if str[j] == ')' {
//					count--
//				}
//				j++
//			}
//			if i+1 == j-1 {
//				continue
//			}
//			data := cal17(str[i+1 : j-1])
//			if data == 0 {
//				return 0
//			}
//			currDepth += data
//			i = j - 1
//		}
//		if str[i] == ')' {
//			maxDepth = max(maxDepth, currDepth)
//			currDepth--
//			continue
//		}
//		if str[i] == '[' {
//			currDepth++
//			count := 1
//			j := i + 1
//			for count > 0 {
//				if j >= len(str) {
//					return 0
//				}
//				if str[j] == '[' {
//					count++
//				} else if str[j] == ']' {
//					count--
//				}
//				j++
//			}
//			if i+1 == j-1 {
//				continue
//			}
//			data := cal17(str[i+1 : j-1])
//			if data == 0 {
//				return 0
//			}
//			currDepth += data
//			i = j - 1
//		}
//		if str[i] == ']' {
//			maxDepth = max(maxDepth, currDepth)
//			currDepth--
//			continue
//		}
//		if str[i] == '{' {
//			currDepth++
//			count := 1
//			j := i + 1
//			for count > 0 {
//				if j >= len(str) {
//					return 0
//				}
//				if str[j] == '{' {
//					count++
//				} else if str[j] == '}' {
//					count--
//				}
//				j++
//			}
//			if i+1 == j-1 {
//				continue
//			}
//			data := cal17(str[i+1 : j-1])
//			if data == 0 {
//				return 0
//			}
//			currDepth += data
//			i = j - 1
//		}
//		if str[i] == '}' {
//			maxDepth = max(maxDepth, currDepth)
//			currDepth--
//			continue
//		}
//	}
//	return maxDepth
//}
