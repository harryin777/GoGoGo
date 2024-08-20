package main

import (
	"fmt"
	"sort"
)

/*
*
现有个任务需要在时间内处理完成，同一时间只能处理一个任务，处理每个任务所需要的时间固定为1。
每个任务都有最晚处理时间限制和报酬，在最晚处理时间点之前处理完成任务才可获得对应的报酬奖励。
可用于处理任务的时间有限，请问在有限的时间内，可获得的最多报酬？
1|1<N<1e9、
输入描述
、1<T<199
第一行输入两个数和，表示个任务和全部任务的最迟的时间节点。
接下来输入N行，每一行输入两个数和L表示一个任务，l<为这个任务的最晚完成时间，
酬。
输出描述
一个整数，表示能够获取的最大报酬。
L为完成该任务能够获得的报

贪心：每个局部最优会导致整体结果最优
*/

func main67() {
	var lastTime, taskCount int
	fmt.Scanln(&lastTime, &taskCount)
	timePointMapTaskVal := make(map[int][]int)
	for i := 0; i < taskCount; i++ {
		var curTaskTimePoint, curTaskVal int
		fmt.Scanf("%d %d", &curTaskTimePoint, &curTaskVal)
		t := min(lastTime, curTaskTimePoint)
		timePointMapTaskVal[t] = append(timePointMapTaskVal[t], curTaskVal)
	}

	ans := 0
	curTaskList := []int{} // 累积的任务列表
	for i := lastTime; i > 0; i-- {
		if tasks, ok := timePointMapTaskVal[i]; ok {
			curTaskList = append(curTaskList, tasks...)
		}
		if len(curTaskList) > 0 {
			sort.Ints(curTaskList)
			ans += curTaskList[len(curTaskList)-1]         // 选出当前最大值
			curTaskList = curTaskList[:len(curTaskList)-1] // 删除已选择的任务
		}

	}
	fmt.Println(ans)
}
