package arrquestions

func BinarySortSearch(input *[]int, item int) int {
	arr := *input
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == item {
			return mid
		}
		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func LinearSearch(input *[]int, item int) int {
	arr := *input
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i
		}
	}

	return -1
}
