package arrquestions_test

import (
	"interviewq/arrquestions"
	"testing"
)

func TestRotateImage(t *testing.T) {
	scenarios := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for _, s := range scenarios {
		result := arrquestions.RotateImage(s.input)

		// compare result
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(result[i]); j++ {
				if result[i][j] != s.expected[i][j] {
					printArray(t, result)
					t.Fatal("Image is not rotated correctly")
				}
			}
		}
	}
}

func BenchmarkRotateImage(b *testing.B) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	for i := 0; i < b.N; i++ {
		arrquestions.RotateImage(input)
	}
}

func TestRotateImageNotOptimizedForSpace(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	result := arrquestions.RotateImageNotOptimizedForSpace(input)

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] != expected[i][j] {
				printArray(t, result)
				t.Fatal("Image is not rotated correctly")
			}
		}
	}
}

func BenchmarkRotateImageNotOptimizedForSpace(b *testing.B) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	for i := 0; i < b.N; i++ {
		arrquestions.RotateImageNotOptimizedForSpace(input)
	}
}

func printArray(t *testing.T, arr [][]int) {
	t.Log("=Result array=\n")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			t.Logf("a[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}
}
