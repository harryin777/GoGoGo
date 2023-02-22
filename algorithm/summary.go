package main

import "math"

// Ten2N 十进制转N进制
func Ten2N(num, N int) int {
	if num == 0 {
		return 0
	}

	slice := make([]int, 0, 100)
	for num != 0 {

		slice = append(slice, num%N)
		num = num / N
	}

	res := 0
	for i := len(slice) - 1; i >= 0; i-- {
		res = res*10 + slice[i]
	}
	return res
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
