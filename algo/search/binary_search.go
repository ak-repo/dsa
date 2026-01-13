package search

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](arr []T, val T) int {
	low := 0
	high := len(arr) - 1

	for high >= low {
		mid := (high + low) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return -1

}
