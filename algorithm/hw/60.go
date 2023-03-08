package main

import (
	"fmt"
	"sort"
)

// https://dream.blog.csdn.net/article/details/129075208
// 任务调度
func main60() {
	var taskList []Task60
	for {
		var id, prio, exe, arr int
		n, _ := fmt.Scan(&id, &prio, &exe, &arr)
		if n <= 0 {
			break
		}
		taskList = append(taskList, Task60{
			Id:       id,
			Priority: prio,
			Execute:  exe,
			Arrive:   arr,
		})
	}

	cal60(taskList)
}

func cal60(task60 []Task60) {
	var cpu Task60
	waiting := make([]Task60, 0, len(task60))
	finish := make([]Task60, 0, len(task60))
	cpu = task60[0]
	currTime := 0
	_ = currTime
	for cpu != (Task60{}) {
		arriveIndex := 1
		if task60[arriveIndex].Arrive < cpu.Arrive+cpu.Execute {
			if task60[arriveIndex].Priority > cpu.Priority {
				cpu.LastTime = task60[arriveIndex].Arrive - cpu.Arrive + cpu.Execute
				waiting = append(waiting, cpu)
				sort.Slice(waiting, func(i, j int) bool {
					if waiting[i].Priority > waiting[j].Priority {
						return true
					} else if waiting[i].Priority == waiting[j].Priority {
						return waiting[i].Arrive < waiting[j].Arrive
					} else {
						return false
					}
				})
				cpu = task60[arriveIndex]
			}
		} else {
			cpu.FinishTime = cpu.Arrive + cpu.Execute
			finish = append(finish, cpu)
			for len(waiting) != 0 {
				if waiting[0].Priority > task60[arriveIndex].Priority {
					cpu = waiting[0]
					waiting = waiting[1:]
				}
			}
			cpu = task60[arriveIndex]
		}
	}

}

type Task60 struct {
	Id         int
	Priority   int
	Execute    int
	Arrive     int
	LastTime   int
	FinishTime int
}
