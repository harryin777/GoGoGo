package sort

func InsertSort(arr []int) (res []int) {
	if len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		var j = 0
		for j = i; j > 0; j-- {
			// 注意这里是和min比较，不是和 j ，虽然第一次j和min相同
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

func Insert2(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		var j = i
		for ; j > 0; j-- {
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
