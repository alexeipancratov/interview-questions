package arrquestions

import "math"

/*
In this solution we leverage the problem condition which states
that the array elements should not be greater than the array length
*/
func GetFirstDuplicate(input *[]int) int {
	arr := *input
	for i := 0; i < len(arr); i++ {
		if arr[abs(arr[i])-1] < 0 {
			return abs(arr[i])
		}

		arr[abs(arr[i])-1] = -arr[abs(arr[i])-1]
	}

	return -1
}

func abs(number int) int {
	return int(math.Abs(float64(number)))
}

// Space complexity can be used to improve time complexity
func GetFirstDuplicateMap(input *[]int) int {
	arr := *input
	visitedElements := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := visitedElements[arr[i]]; ok {
			return arr[i]
		}
		visitedElements[arr[i]] = 1
	}

	return -1
}

func GetFirstDuplicateNSquare(input *[]int) int {
	arr := *input
	minDuplicatePairIndex := len(arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] && j < minDuplicatePairIndex {
				minDuplicatePairIndex = j
			}
		}
	}

	if minDuplicatePairIndex != len(arr) {
		return arr[minDuplicatePairIndex]
	} else {
		return -1
	}
}

func GetFirstDuplicateNonOptimizedMap(input *[]int) int {
	arr := *input
	duplicatePositions := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				duplicatePositions[arr[i]] = j
			}
		}
		if len(duplicatePositions) > 1 {
			keys := make([]int, 0, len(duplicatePositions))
			for k := range duplicatePositions {
				keys = append(keys, k)
			}
			if duplicatePositions[keys[0]] < duplicatePositions[keys[1]] {
				return keys[0]
			} else {
				return keys[1]
			}
		}
	}

	return -1
}
