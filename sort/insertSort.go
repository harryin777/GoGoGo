package sort

func InsertSort(arr []int) (res []int) {
	if len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		var j = 0
		for j = i; j > 0; j-- {
			if arr[j-1] >= min {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = min
	}

	return arr
}

func Insert2(arr []int) (res []int) {
	if len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		j := 0
		for j = i; j > 0; j-- {
			if arr[j-1] >= min {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = min
	}

	return arr
}

func Insert3(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		min := nums[i]
		j := 0
		for j = i; j > 0; j-- {
			if nums[j-1] > min {
				nums[j] = nums[j-1]
			} else {
				break
			}
		}
		nums[j] = min
	}

	return nums
}

func Insert4(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		min := nums[i]
		j := 0
		for j = i; j > 0; j-- {
			if nums[j-1] > min {
				nums[j] = nums[j-1]
			} else {
				break
			}
		}
		nums[j] = min
	}

	return nums
}

func Insert5(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		j := i
		for j = i; j > 0; j-- {
			if nums[j-1] > tmp {
				nums[j] = nums[j-1]
			} else {
				break
			}
		}
		nums[j] = tmp
	}

	return nums
}

func Insert6(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		j := i
		for j = i; j > 0; j-- {
			if nums[j-1] > tmp {
				nums[j] = nums[j-1]
			} else {
				break
			}
		}
		nums[j] = tmp
	}

	return nums
}
