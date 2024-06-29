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
		standard := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[standard] {
				standard = j
			}
		}
		if standard != i {
			arr[i], arr[standard] = arr[standard], arr[i]
		}
	}

	return arr
}
