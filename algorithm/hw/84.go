package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*找座位
 */
func main() {
	var str string
	fmt.Scanln(&str)
	strArr := strings.Split(str, ",")
	intArr := make([]int, 0, len(strArr))
	for _, val := range strArr {
		valInt, _ := strconv.Atoi(val)
		intArr = append(intArr, valInt)
	}

	cal84(intArr)
}

func cal84(flowerbed []int) {
	n := 0
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[i+1] == 0 && flowerbed[i] != 1 {
				n++
				flowerbed[i] = 1
			}
			continue
		}
		if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] != 1 {
				n++
				flowerbed[i-1] = 1
			}
			continue
		}
		if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 && flowerbed[i] != 1 {
			flowerbed[i] = 1
			n++
		}
	}

	fmt.Println(n)
}
