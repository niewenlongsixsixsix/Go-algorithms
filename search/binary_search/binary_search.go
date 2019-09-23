package binary_search

import (
	"github.com/pkg/errors"
	"math"
	"sort"
)

/**
BinarySearch search ordered slice for a specified element. if the element exists, it will return true, false otherwise.
*/

var SliceNotOrderedError = errors.New("the slice isn't ordered")

func BinarySearch(arr []int, specifyElement int) bool {
	arrTemp := sort.IntSlice(arr)
	if !sort.IsSorted(arrTemp) {
		panic(SliceNotOrderedError)
	}
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) / 2
		midElement := arr[int(math.Floor(float64(mid)))]
		if specifyElement == midElement {
			return true
		} else if specifyElement > midElement {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
