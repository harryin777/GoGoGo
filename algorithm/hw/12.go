package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main12() {
	var n int
	fmt.Scan(&n)
	cList := make([][]string, 0, n*2)
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		cList = append(cList, strings.Split(buf.Text(), " "))
	}

	cal12(cList, n)
}

func cal12(cList [][]string, n int) {
	count := 0
	num := 0
	queue := make([]int, 0, len(cList))
	for i := 0; i < len(cList); i++ {
		if len(cList[i]) == 0 {
			continue
		}
		if len(cList[i]) == 1 && cList[i][0] == "remove" && len(queue) != 0 {
			num++
			if num != queue[0] {
				sort.Slice(queue, func(i, j int) bool {
					count++
					return queue[i] < queue[j]
				})
			}
			queue = queue[1:]
			continue
		}

		if cList[i][0] == "head" {
			dataInt, _ := strconv.Atoi(cList[i][2])
			queue = append([]int{dataInt}, queue...)
		} else if cList[i][0] == "tail" {
			dataInt, _ := strconv.Atoi(cList[i][2])
			queue = append(queue, dataInt)
		}
	}

	fmt.Println(count)
}
