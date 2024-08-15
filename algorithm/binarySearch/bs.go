package binarySearch

import (
	"fmt"
	"math"
)

// BS 这种才是二分, 上面那个还是要遍历
func BS(nums []int, target int) int {
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		count++
		fmt.Printf("BS %d \n", count)
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return left
}

// LeftBound
func LeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if left >= right {
		return 0
	}
	// 这里如果加等于有无限循环的可能，因为下面有边界是 right = mid
	for left < right {
		mid := (left + right) >> 1
		// 为什么可以找到左边界，因为在找到target的时候没有立即返回，而是缩小了搜索的右边界
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// 判断 target 是否存在于 nums 中
	// 如果越界，target 肯定不存在，返回 -1
	if left < 0 || left >= len(nums) {
		return -1
	}
	// 判断一下 nums[left] 是不是 target 为什么要加这个判断，因为最上面的循环终止条件是left = right，
	// 而有可能right = len -1 这时候有边界就没有搜索到
	if nums[left] == target {
		return left
	}
	return -1
}

func RightBound(nums []int, target int) int {
	left, right := 0, len(nums)
	if left > right {
		return 0
	}
	for left < right {
		mid := (left + right) >> 1
		// 能够找到右边界的原因是在 mid  <= target 的情况下，缩短了左边界
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
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
