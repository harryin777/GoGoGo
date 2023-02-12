package sort

func BubbleSort(arr []int) (res []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}

	return arr
}
