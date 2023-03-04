package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/129095357
// 内存池
func main43() {
	var caches, req string
	fmt.Scan(&caches, &req)
	data := make([]cache, 0, 10)
	arr := strings.Split(caches, ",")
	for i := 0; i < len(arr); i++ {
		tmp := strings.Split(arr[i], ":")
		t1, _ := strconv.Atoi(tmp[0])
		t2, _ := strconv.Atoi(tmp[1])
		data = append(data, cache{
			Size: t1,
			Num:  t2,
		})
	}
	t3 := strings.Split(req, ",")
	var reqs []int
	for i := 0; i < len(t3); i++ {
		t4, _ := strconv.Atoi(t3[i])
		reqs = append(reqs, t4)
	}

	cal43(data, reqs)
}

func cal43(data []cache, req []int) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Size < data[j].Size
	})
	for i := 0; i < len(req); i++ {
		flag := false
		for j := 0; j < len(data); j++ {
			if data[j].Size >= req[i] && data[j].Num > 0 {
				flag = true
				if i == len(req)-1 {
					fmt.Print("true")
				} else {
					fmt.Print("true,")
				}
				data[j].Num--
				break
			}
		}
		if !flag {
			if i == len(req)-1 {
				fmt.Print("false")
			} else {
				fmt.Print("false,")
			}
		}
	}
}

type cache struct {
	Size int
	Num  int
}
