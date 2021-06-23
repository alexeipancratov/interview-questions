package main

import (
	"fmt"
	"interviewq/memoization"
)

func main() {
	fmt.Printf("Result = %v\n", memoization.Fib(6, nil))  // 8
	fmt.Printf("Result = %v\n", memoization.Fib(7, nil))  // 13
	fmt.Printf("Result = %v\n", memoization.Fib(8, nil))  // 21
	fmt.Printf("Result = %v\n", memoization.Fib(50, nil)) // 12586269025
}
