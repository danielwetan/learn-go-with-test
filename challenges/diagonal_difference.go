package challenges

func diagonalDifference(arr [][]int) int {
	left := 0
	right := 0
	counter := 0

	for i := 0; i < len(arr); i++ {
		left += arr[i][counter]
		right += arr[i][len(arr)-1-counter]
		counter++
	}

	if left > right {
		return left-right
	}

	return right - left
}

