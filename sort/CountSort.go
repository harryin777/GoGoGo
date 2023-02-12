package sort

func CountSort(arr []int) []int {

	max := getMaxOne(arr)
	initArr := make([]int, len(arr))
	copy(initArr, arr)
	countArray := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		countArray[arr[i]]++
	}
	for i := 1; i < len(countArray); i++ {
		countArray[i] = countArray[i] + countArray[i-1]
	}
	for i := 0; i < len(initArr); i++ {
		arr[countArray[initArr[i]]-1] = initArr[i]
	}
	return arr
}

func getMaxOne(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
