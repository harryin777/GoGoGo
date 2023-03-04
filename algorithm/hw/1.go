package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129075208
func main1() {
	var a, b, c, d int
	taskList := make([]Task, 0, 100)
	for {
		n, _ := fmt.Scan(&a, &b, &c, &d)
		if n == 0 {
			break
		}
		taskList = append(taskList, Task{
			ID:          a,
			Prio:        b,
			ExecuteTime: c,
			ArriveTime:  d,
		})
	}
}

func cal(taskList []Task) {
	var currentTask Task
	queue := make([]Task, 0, len(taskList))
	queue = append(queue, taskList[0])
	for len(queue) != 0 {
		currentTask = queue[0]
		queue = queue[1:]
		for i := 1; i < len(taskList); i++ {
			if taskList[i].ArriveTime > currentTask.ArriveTime+currentTask.ExecuteTime {
				fmt.Printf("%v %v", currentTask.ID, currentTask.ArriveTime+currentTask.ExecuteTime)
				queue = append(queue, taskList[i])
				continue
			}
		}
	}
	for i := 0; i < len(taskList); i++ {
		if currentTask == (Task{}) {
			currentTask = taskList[i]
			continue
		}
		if len(queue) != 0 {
			storeTask := queue[len(queue)-1]
			if currentTask.Prio < storeTask.Prio {
				currentTask = storeTask
			}
		}
		if taskList[i].ArriveTime > currentTask.ArriveTime+currentTask.ExecuteTime {
			fmt.Printf("%v %v", currentTask.ID, currentTask.ArriveTime+currentTask.ExecuteTime)
			currentTask = taskList[i]
			continue
		}
		if taskList[i].Prio > currentTask.Prio {
			currentTask.RestTime = taskList[i].ArriveTime - currentTask.ArriveTime
			queue = append(queue, currentTask)
			currentTask = taskList[i]
		} else {
			fmt.Printf("%v %v", currentTask.ID, currentTask.ArriveTime+currentTask.ExecuteTime)
			currentTask = taskList[i]
			continue
		}
	}
}

type Task struct {
	ID          int
	Prio        int
	ExecuteTime int
	ArriveTime  int
	RestTime    int
}
