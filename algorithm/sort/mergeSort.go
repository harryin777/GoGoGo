package sort

// 稳定
func MergeSort2(left, right []int) (res []int) {
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return
}

func MergeSort1(arr []int) (res []int) {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	arrL := MergeSort1(arr[:mid])
	arrR := MergeSort1(arr[mid:])
	return MergeSort2(arrL, arrR)
}

func Merge1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	arrL := Merge1(nums[:mid])
	arrR := Merge1(nums[mid:])

	return Merge2(arrL, arrR)
}

func Merge2(l, r []int) []int {
	i, j := 0, 0
	ans := make([]int, 0, len(l)+len(r))
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			ans = append(ans, l[i])
			i++
		} else {
			ans = append(ans, r[j])
			j++
		}
	}

	ans = append(ans, l[i:]...)
	ans = append(ans, r[j:]...)

	return ans
}
