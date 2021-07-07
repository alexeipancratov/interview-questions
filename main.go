package main

import (
	"fmt"
	"interviewq/dynamicprogramming"
)

func main() {
	fmt.Printf("Result = %v\n", dynamicprogramming.Fib(6, nil))  // 8
	fmt.Printf("Result = %v\n", dynamicprogramming.Fib(7, nil))  // 13
	fmt.Printf("Result = %v\n", dynamicprogramming.Fib(8, nil))  // 21
	fmt.Printf("Result = %v\n", dynamicprogramming.Fib(50, nil)) // 12586269025
}
