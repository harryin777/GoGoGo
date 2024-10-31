package sort

func ShellSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	pace := 0
	for pace < len(arr)/3 {
		pace = pace*3 + 1
	}

	for pace > 0 {
		for i := 0; i < len(arr); i += pace {
			min := arr[i]
			var j = i
			for ; j > 0; j -= pace {
				if arr[j-pace] > min {
					arr[j] = arr[j-pace]
				} else {
					break
				}
			}
			arr[j] = min
		}
		pace = pace / 3
	}

	return arr
}
