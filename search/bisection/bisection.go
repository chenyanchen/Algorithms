package main

import (
	"fmt"
)

const NotExist = -1

func bisection(list []int64, item int64) (int64, bool) {
	var low, mid, high int64
	low, high = 0, int64(len(list)-1)

	for low <= high {
		// the best way to get mid:	mid = low + (high - low) /2
		// 避免值溢出
		mid = int64((low + high) / 2)
		if list[mid] == item {
			return mid, true
		} else if list[mid] >= item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return NotExist, false
}

func main() {
	list := []int64{1, 3, 5, 7, 9}
	var item int64

	fmt.Print("Please enter your number: ")
	fmt.Scanln(&item)

	place, flag := bisection(list, item)
	if !flag {
		fmt.Println("Not exist.")
	} else {
		fmt.Printf("%d in %d place.", item, place)
	}
}
