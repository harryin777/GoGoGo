package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
*
题目描述
疫情期间，需要大家保证一定的社交距离，公司组织开交流会议，座位有一排共个座位，编号分别为N-I]，要求
员工一个接着一个进入会议室，并且可以在任何时候离开会议室。
满足：每当一个员工进入时，需要坐到最大社交距离的座位（例如：位置与左右有员工落座的位置距离分另刂为2和2，
位置与左右有员工落座的位置距离分另刂为2和3，影响因素都为2个位置，则认为座位和与左右位置的社交距离是
一样的）；如果有多个这样的座位，贝刂坐到索引最小的那个座位。
输入描述
会议室座位总数，
（1seatNums599）
员工的进出顺序数组，元素值为1：表示进场；元素值为负数，表示出场（特殊：位置的员工不会离
开），例如一4表示坐在位置4的员工离开（保证有员工坐在该座位上）
输出描述
最后进来员工，他会坐在第几个位置，如果位置已满，则输出一1
*/
func main73() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	var totalSeats int
	fmt.Sscanf(buf.Text(), "%d", &totalSeats)
	buf.Scan()
	s := strings.Split(buf.Text(), " ")
	arr := make([]int, 0, len(s))
	for _, val := range s {
		valInt, _ := strconv.Atoi(val)
		arr = append(arr, valInt)
	}
	cal73(totalSeats, arr)
}

func cal73(totalSearts int, peo []int) {
	s := make([]int, totalSearts)
	s[0] = -1
	lastInIndex := -1
	for i := len(peo) - 1; i >= 0; i++ {
		if peo[i] > 0 {
			lastInIndex = i
			break
		}
	}

	if lastInIndex == 0 {
		fmt.Println(0)
	} else {
		for i := 1; i <= lastInIndex; i++ {
			if i == lastInIndex {
				fmt.Println(updateSearts(s))
				break
			}
			cur := peo[i]
			if cur > 0 {
				curIndex := updateSearts(s)
				if curIndex < 0 {
					continue
				}
				s[curIndex] = -1
			} else {
				s[-cur] = 0
			}
		}
	}

}

func updateSearts(seats []int) int {
	n := len(seats)
	right := n - 1
	for i := n - 1; i >= 0; i-- {
		if seats[i] == -1 {
			right = i
			break
		}
	}

	maxDis := n - 1 - right
	ansIdx := n - 1
	pre := 0
	for i := 1; i <= right; i++ {
		if seats[i] == -1 {
			curDis := (i - pre) / 2
			if curDis > maxDis {
				maxDis = curDis
				ansIdx = pre + curDis
			}
			pre = i
		}
	}
	if maxDis > 0 {
		return ansIdx
	}

	return -1
}

func updateSeatsIn(seats []int) int {
	n := len(seats)
	right := n - 1
	for i := n - 1; i >= 0; i-- {
		if seats[i] == 1 {
			right = i
			break
		}
	}

	ansIdx := n - 1
	maxDis := n - 1 - right
	pre := 0
	for i := 1; i <= right; i++ {
		if seats[i] == 1 {
			curDis := (i - pre) / 2
			if maxDis < curDis {
				maxDis = curDis
				ansIdx = pre + (i-pre)/2
			}
			pre = i
		}
	}
	if maxDis > 0 {
		return ansIdx
	}
	return -1
}

func sss(operations []int, n int) {
	// 初始化座位数组
	seats := make([]int, n)
	seats[0] = 1

	// 找到最后一个正数操作的索引
	lastInOperationIdx := -1
	for i := len(operations) - 1; i >= 0; i-- {
		if operations[i] > 0 {
			lastInOperationIdx = i
			break
		}
	}

	// 如果第一个操作就是正数，输出0
	if lastInOperationIdx == 0 {
		fmt.Println(0)
	} else {
		for i := 1; i <= lastInOperationIdx; i++ {
			if i == lastInOperationIdx {
				ans := updateSeatsIn(seats)
				fmt.Println(ans)
				break
			}
			op := operations[i]
			if op > 0 {
				idx := updateSeatsIn(seats)
				if idx != -1 {
					seats[idx] = 1
				}
			} else {
				idx := -op
				seats[idx] = 0
			}
		}
	}
}
