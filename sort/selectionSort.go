package sort

func SelectionSort(arr []int) (res []int) {

	for i := 0; i < len(arr); i++ {
		standard := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[standard] {
				standard = j
			}
		}
		if standard != i {
			tmp := arr[i]
			arr[i] = arr[standard]
			arr[standard] = tmp
		}
	}

	return arr
}

func Selection2(arr []int) (res []int) {
	if len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			tmp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = tmp
		}
	}

	return arr
}

func Selection3(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		standard := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[standard] {
				standard = j
			}
		}
		if standard != i {
			nums[i], nums[standard] = nums[standard], nums[i]
		}
	}

	return nums
}

func Selection4(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		standard := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[standard] {
				standard = j
			}
		}
		if standard != i {
			nums[i], nums[standard] = nums[standard], nums[i]
		}
	}

	return nums
}

func SelectionSort5(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		standard := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[standard] {
				standard = j
			}
		}
		if standard != i {
			nums[standard], nums[i] = nums[i], nums[standard]
		}
	}

	return nums
}
