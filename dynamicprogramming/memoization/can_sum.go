package memoization

func CanSum(targetSum int, items []int, memo map[int]bool) bool {
	if memo == nil {
		memo = make(map[int]bool)
	}

	if value, ok := memo[targetSum]; ok {
		return value
	} else if targetSum == 0 {
		return true
	} else if targetSum < 0 {
		return false
	} else {
		for i := 0; i < len(items); i++ {
			if CanSum(targetSum-items[i], items, memo) {
				memo[targetSum] = true
				return true
			}
		}

		memo[targetSum] = false
		return false
	}
}

func CanSumNonOptimized(targetSum int, items []int) bool {
	if targetSum == 0 {
		return true
	} else if targetSum < 0 {
		return false
	} else {
		for i := 0; i < len(items); i++ {
			if CanSumNonOptimized(targetSum-items[i], items) {
				return true
			}
		}

		return false
	}
}
