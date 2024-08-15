package sort

func QuickSort(arr []int, left, right int) (res []int) {
	if left > right {
		return arr
	}

	i := left
	j := right
	tmp := arr[i]
	for i < j {
		// 注意这里一定要有等于，因为起始元素tmp和arr[i]是相同的
		// 注意这里的顺序,一定是先 j 后 i
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

func Quick2(nums []int, left, right int) (res []int) {
	if left > right {
		return nums
	}

	i := left
	j := right
	tmp := nums[i]
	for i < j {
		for i < j && nums[j] >= tmp {
			j--
		}
		for i < j && nums[i] <= tmp {
			i++
		}
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left] = nums[i]
	nums[i] = tmp

	Quick2(nums, left, i-1)
	Quick2(nums, i+1, right)
	return nums
}
