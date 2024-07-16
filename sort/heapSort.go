package sort

/*
给定整数数组nums和k，
请返回数组中第k个最大元素，
请注意，你需要找的是数组排序后的第k个最大元素，
而不是第k个不同的元素

假定数组nums的长度为leng

堆的最后一个节点的父节点下标为：leng/2-1

任何一个下标为n的节点的左右子节点下标为：左子节点ln = n*2+1，右子节点rn = n*2+2。
前提是ln和rn小于leng-1,即没有下标溢出，若溢出表明没有该子节点
*/
func swap(a, b *int) {
	*a, *b = *b, *a
}

func HeapSort(arr []int) []int {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	//fmt.Println(arr)

	// 一个个从堆顶取出元素
	for i := n - 1; i > 0; i-- {
		swap(&arr[0], &arr[i]) // 将当前最大值移到数组末尾
		heapify(arr, i, 0)     // 重新堆化，但不包括已排序的元素
	}

	return arr
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		swap(&arr[i], &arr[largest])
		heapify(arr, n, largest)
	}
}

func HeapSort2(arr []int) []int {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heaplify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		swap(&arr[0], &arr[i])
		heaplify(arr, i, 0)
	}

	return arr
}

func heaplify(arr []int, n, i int) {
	max := i
	lson := i*2 + 1
	rson := i*2 + 2
	if lson < n && arr[lson] > arr[max] {
		max = lson
	}
	if rson < n && arr[rson] > arr[max] {
		max = rson
	}
	if max != i {
		swap(&arr[i], &arr[max])
		heapify(arr, n, max)
	}

}
