package binarySearch

import (
	"fmt"
	"math"
)

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

func TripleOne(target float64) {
	var input float64
	input = target

	if input == 0 || input == 1 || input == -1 {
		fmt.Printf("%.1f", input)
		return
	}

	var ne = input < 0
	if ne {
		input = input * -1
	}

	var small = input < 1

	var low, mid, high, triple float64
	if small {
		low = 0
		high = 1
	} else {
		low = 1
		high = input
	}

	for {
		mid = (low + high) / 2
		triple = mid * mid * mid
		if math.Abs(input-triple) < 0.1 {
			break
		} else if triple > input {
			high = mid
		} else if triple < input {
			low = mid
		}
	}

	if ne {
		mid *= -1
	}

	fmt.Printf("%.6f", mid)
}

func NormalBS(nums []int, target int) int {
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		count++
		fmt.Printf("NormalBS %d \n", count)
		mid := (left + right) >> 1
		if nums[mid] > target {
			right -= 1
		} else if nums[mid] < target {
			left += 1
		} else {
			return mid
		}
	}

	return left
}

// BS 这种才是二分, 上面那个还是要遍历
func BS(nums []int, target int) int {
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		count++
		fmt.Printf("BS %d \n", count)
		mid := (left + right) >> 1
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else {
			return mid
		}
	}

	return left
}
