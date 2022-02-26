package challenges

import (
	"sort"
)

/*
Cari bilangan positif terkecil yg tidak ada dalam array

input: [4, 7, 0, 2, 1, 3, 5, 9, 10, 18, -1]
output: 6
*/

func smallestIntNotPresent(arr []int) int {
	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i]+1 != arr[i+1] {
			return arr[i] + 1
		}
	}

	return 0
}

