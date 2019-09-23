package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(BinarySearch(arr, 5))
}
