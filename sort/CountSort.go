package sort

func CountSort(arr []int) []int {

	max := getMaxOne(arr)
	initArr := make([]int, len(arr))
	// copy 函数调用的时候需要提前分配内存和长度,注意是长度不是容量
	copy(initArr, arr)
	countArray := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		countArray[arr[i]]++
	}
	// 注意起始点是 1,后续有 i-1 操作
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

func CountSort2(nums []int) []int {
	backupArr := make([]int, len(nums))
	copy(backupArr, nums)
	maxOne := getMaxOne(nums)
	cumulateArr := make([]int, maxOne+1)

	for n := 0; n < len(nums); n++ {
		cumulateArr[nums[n]]++
	}

	for i := 1; i < len(cumulateArr); i++ {
		cumulateArr[i] = cumulateArr[i] + cumulateArr[i-1]
	}

	for i := 0; i < len(backupArr); i++ {
		nums[cumulateArr[backupArr[i]]-1] = backupArr[i]
	}

	return nums
}

func CountSort3(nums []int) []int {
	initArr := make([]int, len(nums))
	copy(initArr, nums)
	maxVal := getMaxOne(initArr)
	countArr := make([]int, maxVal+1)

	for i := 0; i < len(initArr); i++ {
		countArr[initArr[i]]++
	}

	for i := 1; i < len(countArr); i++ {
		countArr[i] = countArr[i] + countArr[i-1]
	}

	for i := 0; i < len(initArr); i++ {
		nums[countArr[initArr[i]]-1] = initArr[i]
	}

	return nums
}
