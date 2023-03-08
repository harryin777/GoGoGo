package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3",
	4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b",
	12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i",
	19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q",
	27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y",
	35: "z"}

// Ten2N 十进制转N进制
func DecimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 36 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

// Ten2NDecimal 10进制小数转N进制
func Ten2NDecimal(num float64, N int) int {
	if num == 0 {
		return 0
	}

	arr := make([]float64, 0, 100)
	for num != 0 {
		ans := math.Trunc(num * float64(N))
		arr = append(arr, ans)
		num = num*float64(N) - ans
	}

	var res float64
	for i := 0; i < len(arr); i++ {
		res = float64(res*10) + arr[i]
	}
	return int(res)
}

// N210 其他进制转换10进制
func N210(a string, N float64) float64 {
	if len(a) == 0 {
		return 0
	}

	var ans float64
	for index, val := range a {
		t := val - '0'
		if t > 10 {
			val = unicode.ToUpper(val)
			t = val - 55
		}
		ans += float64(t) * math.Pow(N, float64(len(a)-index-1))
	}

	return ans
}

// NNMoveNinetyDegree 把一个n*n的方阵右旋90°
func NNMoveNinetyDegree(mtx [][]int) [][]int {
	if len(mtx) == 0 {
		return [][]int{}
	}

	mtx2 := make([][]int, len(mtx))
	for i := 0; i < len(mtx2); i++ {
		mtx2[i] = make([]int, len(mtx))
	}

	for i := len(mtx) - 1; i >= 0; i-- {
		for j := 0; j < len(mtx); j++ {
			mtx2[j][i] = mtx[len(mtx)-1-i][j]
		}
	}

	return mtx2
}

// SquareNums n 以内的完全平方数
func SquareNums(n int) {
	diff := 3
	square := 1
	for square <= n {
		fmt.Println(square)
		square += diff
		diff += 2
	}
}

func IptoIn(ip string) int {
	ipS := strings.Split(ip, ".")
	res := 0
	for i := 0; i < len(ipS); i++ {
		num, _ := strconv.Atoi(ipS[i])
		res = res*256 + num
	}
	return res
}
