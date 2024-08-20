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
*孙悟空吃蟠桃
 */
func main76() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	arr := strings.Split(buf.Text(), " ")
	peaches := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		val, _ := strconv.Atoi(arr[i])
		peaches = append(peaches, val)
	}
	buf.Scan()
	var h int
	fmt.Scanf("%d", &h)
	cal76(peaches, h)
}

func cal76(peaches []int, h int) {
	if len(peaches) > h {
		fmt.Println(0)
	}
	maxOne := 0
	for i := 0; i < len(peaches); i++ {
		if maxOne < peaches[i] {
			maxOne = peaches[i]
		}
	}

	fmt.Println(1 + sort.Search(maxOne-1, func(speed int) bool {
		speed++
		time := 0
		for i := 0; i < len(peaches); i++ {
			time += (peaches[i] + speed - 1) / speed
		}

		return time <= h
	}))

}
