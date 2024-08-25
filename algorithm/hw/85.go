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
*
最多可以派出多少支团队
*/
func main85() {
	var total int
	fmt.Scanln(&total)
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	strArr := strings.Split(buf.Text(), " ")
	intArr := make([]int, 0, len(strArr))
	for i := 0; i < len(strArr); i++ {
		val, _ := strconv.Atoi(strArr[i])
		intArr = append(intArr, val)
	}
	buf.Scan()
	minC, _ := strconv.Atoi(buf.Text())
	cal85(total, minC, intArr)
}

func cal85(total, minC int, arr []int) {
	sort.Ints(arr)
	ans := 0
	if arr[0] >= minC {
		fmt.Println(len(arr))
		return
	}

	l, r := 0, len(arr)-1
	for l < r {
		if arr[r] >= minC {
			ans++
			r--
		} else if arr[l]+arr[r] >= minC {
			r--
			l++
			ans++
		} else {
			l++
		}
	}

	fmt.Println(ans)
}
