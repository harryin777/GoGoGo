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
题目描述
机器人搬砖，一共有堆砖存放在个不同的仓库中，第堆砖中有bricks[i]块砖头，要求在8小时内搬完。机器人每
小时能搬砖的数量取决于有多少能量格，机器人一个小时中只能在一个仓库中搬砖，机器人的能量格每小时补充一次且能
量格只在这一个小时有效，为使得机器人损耗最小化，尽量减小每次衤卜充的能量格数。
为了保障在8小时内能完成搬砖任务，请计算每小时给机器人充能的最小能量格数。
备注：
1、无需考虑机器人衤卜充能量格的耗时
2、无需考虑机器人搬砖的耗时
3、机器人每小时补充能量格只在这一个小时中有效。
输入描述
程序输入为"391225819"一个整数数组，数组中的每个数字代表第i堆砖的个数，每堆砖的个数不超过1
*/
func main72() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	data := buf.Text()
	arr1 := strings.Split(data, " ")
	arr := make([]int, 0, len(arr1))
	for _, val := range arr1 {
		valInt, _ := strconv.Atoi(val)
		arr = append(arr, valInt)
	}
	cal72(arr)
}

func cal72(arr []int) {
	max := 0

	for _, val := range arr {
		if max < val {
			max = val
		}
	}

	fmt.Println(1 + sort.Search(max-1, func(speed int) bool {
		speed++
		time := 0
		for _, val := range arr {
			time += (val + speed - 1) / speed
		}

		return time <= 8
	}))
}
