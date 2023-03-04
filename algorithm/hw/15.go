package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://dream.blog.csdn.net/article/details/128986346
// 静态扫描最优成本
// 不用整体考虑缓存，对于一个重复的文件来说，如果对于这个文件缓存比全扫描更便宜，就缓存。对整体而言也是更便宜
func main15() {
	var cacheCost int
	fmt.Scan(&cacheCost)
	buf := bufio.NewScanner(os.Stdin)
	var files []file
	var a string
	var b string
	count := 0
	for buf.Scan() {
		str := buf.Text()
		if len(str) == 0 {
			continue
		}
		if count == 0 {
			a = str
		} else {
			b = str
		}
		count++
	}
	arra := strings.Split(a, " ")
	arrb := strings.Split(b, " ")
	for i := 0; i < len(arra); i++ {
		tmp1, _ := strconv.Atoi(arra[i])
		tmp2, _ := strconv.Atoi(arrb[i])
		files = append(files, file{
			Id:   tmp1,
			Size: tmp2,
		})
	}

	cal15(files, cacheCost)
}

func cal15(files []file, cacheCost int) {
	dup := make(map[file]int)
	for _, i2 := range files {
		dup[i2]++
	}
	ans := 0
	for fff, count := range dup {
		total := fff.Size * count
		total = min(total, fff.Size+cacheCost)
		ans += total
	}
	fmt.Println(ans)
}

type file struct {
	Id   int
	Size int
}
