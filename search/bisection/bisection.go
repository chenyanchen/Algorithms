package main

func Bisection(list []int64, item int64) (int64, error) {
	var low, mid, high int64
	low, high = 0, int64(len(list) - 1)
	mid = int64(low + high) / 2

	for low <= high {
		if item == list[mid] {
			return mid, nil
		} else if item >= list[mid] {
			high = mid + 1
		} else {
			low = mid - 1
		}
	}
}