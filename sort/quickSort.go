package sort

func QuickSort(arr []int, left, right int) (res []int) {
	if left > right {
		return arr
	}

	i := left
	j := right
	tmp := arr[i]
	for i < j {
		for i < j && tmp <= arr[j] {
			j--
		}
		for i < j && tmp >= arr[i] {
			i++
		}
		if i < j {
			tmp2 := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp2
		}
	}
	arr[left] = arr[i]
	arr[i] = tmp

	arr = QuickSort(arr, left, i-1)
	arr = QuickSort(arr, i+1, right)
	res = arr
	return res
}
