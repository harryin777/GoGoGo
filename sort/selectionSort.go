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
