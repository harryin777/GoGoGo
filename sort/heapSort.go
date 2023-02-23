package sort

import (
	"fmt"
)

/*
给定整数数组nums和k，
请返回数组中第k个最大元素，
请注意，你需要找的是数组排序后的第k个最大元素，
而不是第k个不同的元素
*/
func swap(a, b *int) {
	*a, *b = *b, *a
}

func HeapSort(nums []int) []int {
	// 堆排序,只能确认第一次个数是最大或最小的
	// 调换第一个元素和最后一个元素位置、从0倒数第二个继续堆排序
	i := len(nums)
	for i > 1 {
		buildHeap(nums, i)
		swap(&nums[0], &nums[i-1])
		i--
	}

	return nums
}

func HeapSort2(nums []int) []int {
	i := len(nums)
	for i > 1 {
		buildHeap(nums, i)
		swap(&nums[0], &nums[i-1])
		i--
	}

	return nums
}

func HeapSort3(nums []int) []int {
	i := len(nums)

	for i > 1 {
		buildHeap(nums, i)
		swap(&nums[0], &nums[i-1])
		i--
	}

	return nums
}

func HeapSort4(nums []int) []int {
	i := len(nums)
	for i > 1 {
		buildHeap4(nums, i)
		swap(&nums[0], &nums[i-1])
		i--
	}

	return nums
}

func buildHeap(nums []int, len int) {
	// 找到最后一个节点的父节点
	parent := len/2 - 1
	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}
	//fmt.Println(nums[0:len])

}

func buildHeap2(nums []int, len int) {
	parent := len/2 - 1
	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}
	fmt.Println(nums[0:len])
}

func buildHeap3(nums []int, len int) {
	parent := len/2 - 1

	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}

	fmt.Println(nums[0:len])
}

func buildHeap4(nums []int, len int) {
	parent := len/2 - 1
	for parent >= 0 {
		heapfiy4(nums, parent, len)
		parent--
	}
}

func heapify(nums []int, parent, len int) {
	// 判断两个子节点是否比父节点大，如果是的话替换
	max := parent
	lson := parent*2 + 1
	rson := parent*2 + 2
	if lson < len && nums[lson] > nums[max] {
		// 左节点是否大于父节点
		max = lson
	}
	if rson < len && nums[rson] > nums[max] {
		// 右节点是否大于父节点
		max = rson
	}
	if parent != max {
		swap(&nums[max], &nums[parent])
		heapify(nums, max, len)
	}
}

func heapify2(nums []int, parent, len int) {
	max := parent
	lson := parent*2 + 1
	rson := parent*2 + 2
	if lson < len && nums[lson] > nums[max] {
		max = lson
	}
	if rson < len && nums[rson] > nums[max] {
		max = rson
	}
	if parent != max {
		swap(&nums[max], &nums[parent])
		heapify2(nums, max, len)
	}
}

func heapify3(nums []int, parent, len int) {
	max := parent
	lson := parent*2 + 1
	rson := parent*2 + 2
	for lson < len && nums[lson] > nums[max] {
		max = lson
	}
	for rson < len && nums[rson] > nums[max] {
		max = rson
	}
	if parent != max {
		swap(&nums[max], &nums[parent])
		heapify3(nums, max, len)
	}
}

func heapfiy4(nums []int, parent, len int) {
	max := parent
	lson := parent*2 + 1
	rson := parent*2 + 2
	for lson < len && nums[lson] > nums[max] {
		max = lson
	}
	for rson < len && nums[rson] > nums[max] {
		max = rson
	}
	if parent != max {
		swap(&nums[max], &nums[parent])
		heapfiy4(nums, max, len)
	}
}
