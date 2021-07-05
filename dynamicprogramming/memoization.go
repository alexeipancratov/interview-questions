package dynamicprogramming

func Fib(n int, memo map[int]int) int {
	if memo == nil {
		memo = make(map[int]int)
	}

	if val, ok := memo[n]; ok {
		return val
	}

	if n <= 2 {
		return 1
	}

	memo[n] = Fib(n-1, memo) + Fib(n-2, memo)

	return memo[n]
}

func FibNonOptimized(n int) int {
	if n <= 2 {
		return 1
	}

	return FibNonOptimized(n-1) + FibNonOptimized(n-2)
}
