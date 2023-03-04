package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main4() {
	var a int
	fmt.Scan(&a)
	money := make([]string, 0, a)
	for i := 0; i < a; i++ {
		var tmp string
		fmt.Scan(&tmp)
		money = append(money, tmp)
	}
	cal4(money)
}

func cal4(money []string) {
	m := map[string]float64{
		"CNY": 100,
		"JPY": 1825,
		"HKD": 123,
		"EUR": 14,
		"GBP": 12,
	}

	var ans float64
	for i := 0; i < len(money); i++ {
		var pre, after float64
		preIndex := 0
		flag := ""
		if index := strings.Index(money[i], "CNY"); index != -1 {
			flag = "c"
			preIndex = index + 3
			pre, _ = strconv.ParseFloat(money[i][0:index], 64)
		}
		if index := strings.Index(money[i], "fen"); index != -1 {
			flag = "c"
			after, _ = strconv.ParseFloat(money[i][preIndex:index], 64)
		}
		if index := strings.Index(money[i], "JPY"); index != -1 {
			flag = "j"
			preIndex = index + 3
			pre, _ = strconv.ParseFloat(money[i][0:index], 64)
		}
		if index := strings.Index(money[i], "sen"); index != -1 {
			flag = "j"
			after, _ = strconv.ParseFloat(money[i][preIndex:index], 64)
		}
		if index := strings.Index(money[i], "HKD"); index != -1 {
			flag = "h"
			preIndex = index + 3
			pre, _ = strconv.ParseFloat(money[i][0:index], 64)
		}
		if index := strings.Index(money[i], "cents"); index != -1 {
			flag = "h"
			after, _ = strconv.ParseFloat(money[i][preIndex:index], 64)
		}
		if index := strings.Index(money[i], "EUR"); index != -1 {
			flag = "e"
			preIndex = index + 3
			pre, _ = strconv.ParseFloat(money[i][0:index], 64)
		}
		if index := strings.Index(money[i], "eurocents"); index != -1 {
			flag = "e"
			after, _ = strconv.ParseFloat(money[i][preIndex:index], 64)
		}
		if index := strings.Index(money[i], "GBP"); index != -1 {
			flag = "g"
			preIndex = index + 3
			pre, _ = strconv.ParseFloat(money[i][0:index], 64)
		}
		if index := strings.Index(money[i], "pence"); index != -1 {
			flag = "g"
			after, _ = strconv.ParseFloat(money[i][preIndex:index], 64)
		}

		switch flag {
		case "c":
			if pre != 0 {
				pre = pre * 100
			}
			ans += pre + after
		case "j":
			if pre != 0 {
				pre = pre / m["JPY"] * 100 * 100
			}
			if after != 0 {
				after = after / m["JPY"] * 100
			}
			ans += pre + after
		case "h":
			if pre != 0 {
				pre = pre / m["HKD"] * 100 * 100
			}
			if after != 0 {
				after = after / m["HKD"] * 100
			}
			ans += pre + after
		case "e":
			if pre != 0 {
				pre = pre / m["EUR"] * 100 * 100
			}
			if after != 0 {
				after = after / m["EUR"] * 100
			}
			ans += pre + after
		case "G":
			if pre != 0 {
				pre = pre / m["GBP"] * 100 * 100
			}
			if after != 0 {
				after = after / m["GBP"] * 100
			}
			ans += pre + after
		}

	}
	fmt.Printf("%.0f", math.Trunc(ans))
}
