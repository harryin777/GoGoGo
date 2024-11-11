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

func Merge1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	arrL := Merge1(arr[:mid])
	arrR := Merge1(arr[mid:])
	return Merge2(arrL, arrR)
}

func Merge2(arr1, arr2 []int) []int {
	i, j := 0, 0
	ans := make([]int, 0, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			ans = append(ans, arr1[i])
			i++
		} else {
			ans = append(ans, arr2[j])
			j++
		}
	}
	ans = append(ans, arr1[i:]...)
	ans = append(ans, arr2[j:]...)

	return ans
}
