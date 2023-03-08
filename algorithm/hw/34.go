package main

import (
	"fmt"
	"strconv"
)

// https://dream.blog.csdn.net/article/details/129004926
// 数据分类
func main34() {
	var c, b int
	fmt.Scan(&c, &b)
	var data []int
	for i := 0; i < 10; i++ {
		var tmp int
		fmt.Scan(&tmp)
		data = append(data, tmp)
	}
	call34(c, b, data)
}

func call34(c, b int, data []int) {
	dul := make(map[int]int)
	for i := 0; i < len(data); i++ {
		r := byteSum(data[i]) % b
		if r < c {
			dul[r]++
		}
	}
	ans := 0
	for _, val := range dul {
		ans = max(ans, val)
	}

	fmt.Println(ans)
}

func byteSum(a int) int {
	aStr := strconv.FormatInt(int64(a), 10)
	ans := 0
	for i := 0; i < len(aStr); i++ {
		// 0xff 代表一个255，低8位全是1，其他位全是0
		// &操作相当于把这个数字的最后一个字节8个比特位的数字得出来了
		ans += a >> (i * 8) & 0xff
	}
	return ans
}
