package sort

import (
	"fmt"
	"math"
	"math/rand"
)

var arr []int
var arrChan []chan chan int
var exp int = 12 //sort 2**exp integers using exp goroutines

func MultiMergeSort() {
	initArr(int(math.Pow(2, float64(exp))))

	for i := 1; i < exp; i++ {
		go work(i)
	}
	boot(0)

	result := <-arrChan[len(arrChan)-1]
	for i := range result {
		fmt.Println(i)
	}
}

func initArr(num int) {
	arr = make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = rand.Intn(10000)
	}

	arrChan = make([]chan chan int, exp+1)
	for i := 0; i < exp+1; i++ {
		arrChan[i] = make(chan chan int, 2)
	}
}

func boot(index int) {
	for i := 0; i < len(arr); i += 2 {
		ch := make(chan int, int64(math.Pow(2, float64(index+1))))
		arrChan[index+1] <- ch
		min := i
		max := i + 1
		if arr[min] > arr[max] {
			max = i
			min = i + 1
		}
		ch <- arr[min]
		ch <- arr[max]
		close(ch)
	}
}

func work(index int) {
	//fmt.Println("work:", index)
	in := arrChan[index]
	for {
		ch1, ok := <-in
		if !ok {
			close(arrChan[index+1])
			break
		}
		ch2 := <-in

		ch := make(chan int, int64(math.Pow(2, float64(index+1))))
		arrChan[index+1] <- ch
		var v1, v2 int
		var v1min, init, v1closed, v2closed bool
		for {
			if !init {
				init = true
				v1 = <-ch1
				v2 = <-ch2
				if v1 < v2 {
					v1min = true
				}
			} else {
				if v1min {
					v1, ok = <-ch1
					if !ok {
						v1closed = true
						v1min = false
					} else if !v2closed && v1 > v2 {
						v1min = false
					}
				} else {
					v2, ok = <-ch2
					if !ok {
						v2closed = true
						v1min = true
					} else if !v1closed && v1 < v2 {
						v1min = true
					}
				}
			}
			if v1closed && v2closed {
				close(ch)
				break
			}
			if v1min {
				ch <- v1
			} else {
				ch <- v2
			}
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	n := len(lists) / 2
	return merge(mergeKLists(lists[:n]), mergeKLists(lists[n:]))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	node := head

	for {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				node.Next = l1
				node, l1 = node.Next, l1.Next
			} else {
				node.Next = l2
				node, l2 = node.Next, l2.Next
			}
		} else {
			if l1 == nil {
				node.Next = l2
			} else {
				node.Next = l1
			}
			break
		}
	}
	return head.Next
}
