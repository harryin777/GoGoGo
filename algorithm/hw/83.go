package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
多段线数据压缩
*/

func main83() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	strArr := strings.Split(buf.Text(), " ")
	intArr := make([]int, 0, len(strArr))
	for i := 0; i < len(strArr); i++ {
		val, _ := strconv.Atoi(strArr[i])
		intArr = append(intArr, val)
	}
	cal83(intArr)
}

func cal83(arr []int) {
	ans := make([]int, 0, len(arr))
	ans = append(ans, arr[0], arr[1])
	nx, ny := 0, 0
	for i := 2; i < len(arr)-1; i = i + 2 {
		if nx == 0 && ny == 0 {
			nx, ny = arr[i]-arr[i-2], arr[i+1]-arr[i-1]
			continue
		}
		curx, cury := arr[i]-arr[i-2], arr[i+1]-arr[i-1]
		if curx == nx && cury == ny {
			continue
		} else {
			ans = append(ans, arr[i-2], arr[i-1])
			nx, ny = curx, cury
		}
	}
	ans = append(ans, arr[len(arr)-2], arr[len(arr)-1])
	for i := 0; i < len(ans); i++ {
		fmt.Printf("%v ", ans[i])
	}
}
