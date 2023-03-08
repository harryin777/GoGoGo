package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func main14() {
	var N float64
	var sub1, sub2 string
	fmt.Scan(&N, &sub1, &sub2)
	cal14(N, sub1, sub2)
}

func cal14(n float64, s1, s2 string) {
	s1Int := N210(n, s1)
	s2Int := N210(n, s2)
	fmt.Println(Ten2N(int(n), strconv.Itoa(s1Int-s2Int)))
}

func N210(n float64, str string) int {
	ans := 0
	for i := 0; i < len(str); i++ {
		t := str[i]
		t = t - '0'
		if t > 10 {
			t = uint8(unicode.ToUpper(rune(t)))
			t = t - 55
		}
		ans += int(t) * int(math.Pow(n, float64(len(str)-i-1)))
	}
	return ans
}

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3",
	4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b",
	12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i",
	19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q",
	27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y",
	35: "z"}

func Ten2N(n int, str string) string {
	var ans string
	var remainderStr string
	num, _ := strconv.Atoi(str)
	for num != 0 {
		remainder := num % n

		if remainder > 9 && remainder < 36 {
			remainderStr = tenToAny[remainder]
		} else {
			remainderStr = strconv.Itoa(remainder)
		}
		ans = remainderStr + ans
		num /= n
	}
	return ans
}
