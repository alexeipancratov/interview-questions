package memoization

import "fmt"

/*
Calculate number of ways to travel from the top left to the bottom right point.
*/
func GridTraveler(n int, m int, memo map[string]int) int {
	if memo == nil {
		memo = make(map[string]int)
	}

	if n == 1 && m == 1 {
		return 1
	} else if n == 0 || m == 0 {
		return 0
	}

	str := fmt.Sprintf("%v, %v", n, m)

	if value, ok := memo[str]; ok {
		return value
	} else {
		memo[str] = GridTraveler(n-1, m, memo) + GridTraveler(n, m-1, memo)

		return memo[str]
	}
}

func GridTravelerNonOptimized(n int, m int) int {
	if n == 1 && m == 1 {
		return 1
	} else if n == 0 || m == 0 {
		return 0
	}

	return GridTravelerNonOptimized(n-1, m) + GridTravelerNonOptimized(n, m-1)
}
