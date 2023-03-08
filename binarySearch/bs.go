package binarySearch

//
//

func LeftBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func RightBound(arr []int, target int) int {
	left, right := 0, len(arr)
	if left > right {
		return 0
	}
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}
