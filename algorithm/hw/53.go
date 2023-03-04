package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128985909
// 最多等和不相交连续子序列
func main() {
	var count int
	fmt.Scan(&count)
	data := make([]int, 0, count)
	for i := 0; i < count; i++ {
		var tmp int
		fmt.Scan(&tmp)
		data = append(data, tmp)
	}

	cal53(data)
}

func cal53(data []int) {
	dupCount := make(map[int]string)
	dup := make(map[int]int)
	mlen := 0
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			tmp := sum53(data[i : j+1])
			flag := false
			for k := i; k <= j; k++ {
				if strings.Contains(dupCount[tmp], strconv.FormatInt(int64(k), 10)) {
					flag = true
					break
				}
			}
			if !flag {
				for k := i; k <= j; k++ {
					dupCount[tmp] = dupCount[tmp] + strconv.FormatInt(int64(k), 10)
				}
				dup[tmp]++
			}

		}
	}
	for key, val := range dup {
		if key == 0 {
			continue
		}
		mlen = max(mlen, val)
	}

	fmt.Println(mlen)
}

func sum53(data []int) int {
	ans := 0
	for i := 0; i < len(data); i++ {
		ans += data[i]
	}

	return ans
}
