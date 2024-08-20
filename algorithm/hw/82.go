package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
查找接口成功率最优时间段
2
0 0 100 2 2 99 0 2 111 4 0
*/

func main82() {
	var minEvg int
	fmt.Scanln(&minEvg)
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	str := strings.Split(buf.Text(), " ")
	sInt := make([]int, 0, len(str))
	for i := 0; i < len(str); i++ {
		val, _ := strconv.Atoi(str[i])
		sInt = append(sInt, val)
	}
	cal82(minEvg, sInt)
}

func cal82(minEvg int, arr []int) {
	n := len(arr)
	res := make([][]int, 0, 10)
	for i := n; i >= 2; i-- {
		threshold := i * minEvg
		sum := 0
		for j := 0; j < i; j++ {
			sum += arr[j]
		}
		if sum <= threshold {
			res = append(res, []int{0, i - 1})
		}
		for j := 0; j < n-i; j++ {
			sum += arr[j+i]
			sum -= arr[j]
			if sum <= threshold {
				res = append(res, []int{j + 1, j + i})
			}
		}
	}

	if len(res) != 0 {
		standard := 0
		for i := 0; i < len(res); i++ {
			if standard == 0 {
				standard = res[i][1] - res[i][0]
			}
			if res[i][1]-res[i][0] >= standard {
				fmt.Printf("%v-%v ", res[i][0], res[i][1])
			} else {
				break
			}
		}
	}
}
