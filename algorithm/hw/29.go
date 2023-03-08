package main

import (
	"fmt"
)

// https://dream.blog.csdn.net/article/details/128985679
// 最小传递延迟
func main29() {
	var nodeCount, listLength int
	fmt.Scan(&nodeCount, &listLength)
	var tuList []Tu
	for i := 0; i < listLength; i++ {
		var source, target, delay int
		fmt.Scan(&source, &target, &delay)
		tuList = append(tuList, Tu{
			Source: source,
			Target: target,
			Delay:  delay,
		})
	}
	var sNode, tNode int
	fmt.Scan(&sNode, &tNode)
	cal29(tuList, sNode, tNode)
}

func cal29(tuList []Tu, sNode, tNode int) {
	// 存储每一个节点，如果节点可以相连也存储
	grid := make([][]Tu, 0, len(tuList))
	// 存储每一个链表的长度
	pathLength := make(map[int]int)
	for i := 0; i < len(tuList); i++ {
		for j := 0; j < len(grid); j++ {
			if _, e := pathLength[j]; !e {
				continue
			}
			if tuList[i].Source == grid[j][pathLength[j]-1].Target {
				grid[j] = append(grid[j], tuList[i])
				pathLength[j]++
			}
		}
		grid = append(grid, []Tu{tuList[i]})
		pathLength[len(grid)-1]++
	}
	minDelay := int(^uint(0) >> 1)
	for i := 0; i < len(grid); i++ {
		currDelay := 0
		for j := 0; j < len(grid[i]); j++ {
			if j == 0 && grid[i][j].Source != sNode {
				break
			}
			currDelay += grid[i][j].Delay
			if grid[i][j].Target == tNode {
				break
			}
		}
		if currDelay != 0 {
			minDelay = min(minDelay, currDelay)
		}
	}
	fmt.Println(minDelay)
}

type Tu struct {
	Source int
	Target int
	Delay  int
}
