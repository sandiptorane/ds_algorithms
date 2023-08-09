package search

func BinarySearch(arr []int, low, high, key int) bool {
	if high >= low {
		// find middle
		mid := (low + high) / 2

		// If element is present at the middle itself
		if arr[mid] == key {
			return true
		}

		// If element is greater than mid, then it
		// can only be present in right subarray else in left subarray
		if arr[mid] > key {
			return BinarySearch(arr, low, mid-1, key)
		} else {
			return BinarySearch(arr, mid+1, high, key)
		}
	}

	return false
}
