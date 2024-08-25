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
*身高体重排序
 */
func main87() {
	var total int
	fmt.Scanln(&total)
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	heiStr := strings.Split(buf.Text(), " ")
	heiInt := make([]int, 0, len(heiStr))
	for i := 0; i < len(heiStr); i++ {
		val, _ := strconv.Atoi(heiStr[i])
		heiInt = append(heiInt, val)
	}
	buf.Scan()
	weiStr := strings.Split(buf.Text(), " ")
	weiInt := make([]int, 0, len(weiStr))
	for i := 0; i < len(weiStr); i++ {
		val, _ := strconv.Atoi(weiStr[i])
		weiInt = append(weiInt, val)
	}
	cal87(total, heiInt, weiInt)
}

type s struct {
	height int
	weight int
	no     int
}

func cal87(total int, hArr, wArr []int) {
	sSlice := make([]s, 0, total)
	for i := 0; i < total; i++ {
		sSlice = append(sSlice, s{
			height: hArr[i],
			weight: wArr[i],
			no:     i + 1,
		})
	}

	sort.Slice(sSlice, func(i, j int) bool {
		if sSlice[i].height < sSlice[j].height {
			return true
		} else if sSlice[i].height == sSlice[j].height {
			if sSlice[i].weight < sSlice[j].weight {
				return true
			} else {
				return false
			}
		}

		return false
	})
	for i := 0; i < len(sSlice); i++ {
		fmt.Printf("%v", sSlice[i].no)
	}

}
