package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128985421
// 简易内存池
func main28() {
	var commandCount int
	fmt.Scan(&commandCount)
	var command []string
	for i := 0; i < commandCount; i++ {
		var tmp string
		fmt.Scan(&tmp)
		command = append(command, tmp)
	}

	cal28(command)
}

func cal28(command []string) {
	// 用一个长度100的slice去当内存
	store := make([]int, 100)
	// 存储每次分配的长度，以首地址为标记
	indexLenMap := make(map[int]int)
	for i := 0; i < len(command); i++ {
		parts := strings.Split(command[i], "=")
		if parts[0] == "REQUEST" {
			flag := true
			for j := 0; j < len(store); j++ {
				if store[j] != 0 {
					continue
				}
				lenInt, _ := strconv.Atoi(parts[1])
				if j+lenInt >= len(store) {
					fmt.Println("error")
					return
				} else if lenInt == 0 {
					fmt.Println("error")
					return
				}
				flag = false
				next := false
				// 这一部分判断空闲的地址够不够本次申请的大小
				for p := j; p < j+lenInt; p++ {
					if store[p] == 1 {
						// 不够，把已经置1的还原
						for m := j; m < p; m++ {
							store[m] = 0
						}
						next = true
						// 找到当前这块分配的内存的尾部
						j = indexLenMap[p] - 1
						break
					}
					store[p] = 1
				}
				if next {
					continue
				}
				indexLenMap[j] = lenInt
				fmt.Println(j)
				break
			}
			if flag {
				fmt.Println("error")
				return
			}
		} else if parts[0] == "RELEASE" {
			pIndex, _ := strconv.Atoi(parts[1])
			if _, e := indexLenMap[pIndex]; !e {
				fmt.Println("error")
				return
			}
			for j := pIndex; j < indexLenMap[pIndex]; j++ {
				store[j] = 0
			}
		}

	}
}
