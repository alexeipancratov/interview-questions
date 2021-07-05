package dynamicprogramming_test

import (
	"interviewq/dynamicprogramming"
	"testing"
)

func benchmarkFib(n int, b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dynamicprogramming.Fib(n, nil)
	}
}

func BenchmarkFib_10(b *testing.B) {
	benchmarkFib(10, b)
}

func BenchmarkFib_20(b *testing.B) {
	benchmarkFib(20, b)
}

func benchmarkFibNonOptimized(n int, b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dynamicprogramming.FibNonOptimized(n)
	}
}

func BenchmarkFibNonOptimized_10(b *testing.B) {
	benchmarkFibNonOptimized(10, b)
}

func BenchmarkFibNonOptimized_20(b *testing.B) {
	benchmarkFibNonOptimized(20, b)
}
