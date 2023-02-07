package sort

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
