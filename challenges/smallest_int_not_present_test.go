package challenges

import (
	"reflect"
	"testing"
)

func TestSmallestIntNotPresent(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		input := []int{4, 7, 0, 2, 1, 3, 5, 9, 10, 18, -1}
		got := smallestIntNotPresent(input)
		want := 6

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
