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
