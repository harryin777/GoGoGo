package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
* 回收银饰
 */
func main75() {
	var totalCount int
	fmt.Scanln(&totalCount)
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	arr := buf.Text()
	arrStr := strings.Split(arr, " ")
	slivers := make([]int, 0, len(arrStr))
	for i := 0; i < len(arrStr); i++ {
		sliver, _ := strconv.Atoi(arrStr[i])
		slivers = append(slivers, sliver)
	}
	cal75(slivers)
}

func cal75(slivers []int) {
here:
	sort.Slice(slivers, func(i, j int) bool {
		if slivers[i] < slivers[j] {
			return false
		} else {
			return true
		}
	})

	for len(slivers) >= 3 {
		z, y, x := slivers[0], slivers[1], slivers[2]
		if z == y && y == x {
			slivers = slivers[3:]
		} else if z != y && y == x {
			slivers = slivers[3:]
			slivers = append(slivers, z-y)
			goto here
		} else if z == y && y != x {
			slivers = slivers[3:]
			slivers = append(slivers, y-x)
			goto here
		} else if z != y && y != x {
			slivers = slivers[3:]
			ans := (z - y) - (y - x)
			slivers = append(slivers, abs(ans))
			goto here
		}
	}

	if len(slivers) == 0 {
		fmt.Println(0)
	} else if len(slivers) == 1 {
		fmt.Println(slivers[0])
	} else {
		if slivers[0] > slivers[1] {
			fmt.Println(slivers[0])
		} else {
			fmt.Println(slivers[1])
		}
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x

}
