package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128989893
// 获取软件最大版本号
func main18() {
	var versions []string
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		versions = append(versions, buf.Text())
	}
	cal18(versions)
}

func cal18(versions []string) {
	var v1 []int
	var v2 []int
	var v1c string
	var v2c string
	for i := 0; i < len(versions); i++ {
		data1 := strings.Split(versions[i], ".")
		for j := 0; j < len(data1); j++ {
			tmp, _ := strconv.Atoi(data1[j])
			if i == 0 {
				v1 = append(v1, tmp)
			} else if i == 1 {
				v2 = append(v2, tmp)
			}
		}
		if len(data1) > 2 {
			data2 := strings.Split(data1[2], "-")
			if len(data2) > 1 {
				tmp, _ := strconv.Atoi(data2[0])
				if i == 0 {
					v1c = data2[1]
					v1 = append(v1, tmp)
				} else if i == 1 {
					v2c = data2[1]
					v2 = append(v2, tmp)
				}
			}
		}
	}

	maxLenPre := min(len(v1), len(v2))
	for i := 0; i < maxLenPre; i++ {
		if v1[i] > v2[i] {
			fmt.Println(versions[0])
			return
		} else if v1[i] < v2[i] {
			fmt.Println(versions[1])
			return
		}
	}
	if len(v2) > len(v1) {
		fmt.Println(versions[1])
		return
	} else if len(v2) < len(v1) {
		fmt.Println(versions[0])
		return
	}

	maxCl := min(len(v1c), len(v2c))
	for i := 0; i < maxCl; i++ {
		if v1c[i] > v2c[i] {
			fmt.Println(versions[0])
			return
		} else if v1c[i] < v2c[i] {
			fmt.Println(versions[1])
			return
		}
	}
	if len(v2c) > len(v1c) {
		fmt.Println(versions[1])
		return
	} else if len(v2c) < len(v1c) {
		fmt.Println(versions[0])
		return
	}

	fmt.Println(versions[0])
}
