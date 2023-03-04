package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128989908
// 服务依赖

func main19() {
	buf := bufio.NewScanner(os.Stdin)
	count := 0
	var arr []string
	var dis []string
	for buf.Scan() {
		str := buf.Text()
		if count == 0 {
			arr = strings.Split(str, ",")
		} else {
			dis = strings.Split(str, ",")
		}
		count++
	}

	cal19(arr, dis)
}

func cal19(arr, dis []string) {
	relation := make([][]string, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		data := strings.Split(arr[i], "-")
		flag := true
		for j := 0; j < len(relation); j++ {
			if relation[j][0] == data[0] {
				relation[j] = append([]string{data[1]}, relation[j]...)
				flag = false
				break
			}
		}
		if flag {
			tmp := []string{data[1], data[0]}
			relation = append(relation, tmp)
		}
	}

	ans := make([]string, 0, len(arr))
	for i := 0; i < len(relation); i++ {
		for j := 0; j < len(relation[i]); j++ {
			for p := 0; p < len(dis); p++ {
				if relation[i][j] == dis[p] {
					goto out
				}
			}
			ans = append(ans, relation[i][j])
		}
	out:
	}
	if len(ans) == 0 {
		fmt.Println(",")
	} else {
		for i := 0; i < len(ans); i++ {
			if i != len(ans)-1 {
				fmt.Printf("%v,", ans[i])
			} else {
				fmt.Printf("%v", ans[i])
			}
		}
	}
}
