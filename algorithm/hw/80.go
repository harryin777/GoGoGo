package main

import "fmt"

/**
找数字
*/

func main80() {
	var n int
	fmt.Scanln(&n)
	cal80(n)
}

func cal80(n int) {
	cnt0, cnt1 := 0, 0
	ans := n
	for (n & 1) == 0 {
		cnt0++
		n >>= 1
	}

	for (n & 1) == 1 {
		cnt1++
		n >>= 1
	}

	k := cnt0 + cnt1
	ans >>= k
	ans <<= k
	mask := 1 << k
	ans |= mask

	k = cnt1 - 1
	mask = (1 << k) - 1
	ans |= mask

	fmt.Println(ans)
}
