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
// 不稳定
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

func heapSort2(arr []int) []int {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify2(arr, len(arr), i)
	}

	for i := len(arr) - 1; i > 0; i-- {
		swap(&arr[i], &arr[0])
		heapify2(arr, i, 0)
	}

	return arr
}

func heapify2(arr []int, n, i int) []int {
	largest := i
	lson := i*2 + 1
	rson := i*2 + 2
	if lson < n && arr[lson] > arr[largest] {
		largest = lson
	}
	if rson < n && arr[rson] > arr[largest] {
		largest = rson
	}
	if largest != i {
		swap(&arr[largest], &arr[i])
		heapify2(arr, n, largest)
	}

	return arr
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	hp := &heapSort{
		size: k,
		arr:  []int{},
	}

	for i := 0; i < len(nums); i++ {
		hp.add(nums[i])
	}
	return hp.arr[0]
}

type heapSort struct {
	arr  []int
	size int
}

func (hp *heapSort) add(num int) {
	if len(hp.arr) < hp.size {
		hp.arr = append(hp.arr, num)
		for i := len(hp.arr) - 1; i > 0; {
			parent := (i - 1) / 2
			if parent >= 0 && hp.arr[parent] > hp.arr[i] {
				hp.swap(parent, i)
				i = parent
			} else {
				break
			}
		}
	} else if num > hp.arr[0] {
		hp.arr[0] = num
		hp.heapify(0)
	}
}

func (hp *heapSort) heapify(i int) {
	max := i
	lson := i*2 + 1
	rson := i*2 + 2
	n := len(hp.arr)
	for lson < n && hp.arr[lson] < hp.arr[max] {
		max = lson
	}
	for rson < n && hp.arr[rson] < hp.arr[max] {
		max = rson
	}
	if max != i {
		hp.swap(i, max)
		hp.heapify(max)
	}
}

func (hp *heapSort) swap(i, j int) {
	hp.arr[i], hp.arr[j] = hp.arr[j], hp.arr[i]
}
