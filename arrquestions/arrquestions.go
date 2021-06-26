package arrquestions

func RotateImage(image [][]int) [][]int {
	arrSize := len(image)
	// transpose image
	for i := 0; i < arrSize; i++ {
		for j := i; j < arrSize; j++ {
			temp := image[i][j]
			image[i][j] = image[j][i]
			image[j][i] = temp
		}
	}

	// reverse each row
	for i := 0; i < arrSize; i++ {
		for j := 0; j < arrSize/2; j++ {
			temp := image[i][j]
			image[i][j] = image[i][arrSize-j-1]
			image[i][arrSize-j-1] = temp
		}
	}

	return image
}

func RotateImageNotOptimizedForSpace(image [][]int) [][]int {
	result := make([][]int, len(image))
	for i := range image {
		result[i] = make([]int, len(image[i]))
	}

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image); j++ {
			row := j
			col := len(image) - i - 1
			result[row][col] = image[i][j]
		}
	}
	return result
}
