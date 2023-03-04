package main

import (
	"fmt"
	"sort"
)

// https://blog.csdn.net/hihell/article/details/128989948
// 字母计数
func main7() {
	var str string
	fmt.Scan(&str)

	cal7(str)
}

func cal7(str string) {
	list := make([]C, 0, len(str))
	mapC := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		mapC[str[i]]++
	}
	for b, i := range mapC {
		list = append(list, C{
			C:     b,
			Count: i,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Count > list[j].Count {
			return true
		} else if list[i].Count == list[j].Count {
			return list[i].C > list[i].C
		} else {
			return false
		}
	})

	for i := 0; i < len(list); i++ {
		fmt.Printf("%v:%d;", string(list[i].C), list[i].Count)
	}
}

type C struct {
	C     byte
	Count int
}
