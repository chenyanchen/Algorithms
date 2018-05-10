/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

package main

import (
	"fmt"
)

const NotExist = -1

func main() {
	list := []int{1, 3, 5, 7, 9, 11, 13}
	var item int

	fmt.Print("Please enter your number: ")
	fmt.Scanln(&item)

	place := F(list, item)
	fmt.Println(place)
}

var low, mid, high int

func bisection(list []int, item int) (int, bool) {
	var low, mid, high int
	low, high = 0, int(len(list)-1)

	for low <= high {
		// the best way to get mid:	mid = low + (high - low) /2
		// 避免值溢出
		mid = int((low + high) / 2)
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

// 二分查找法递归实现方式
func F(l []int, item int) int {
	low, high = 0, len(l)-1

	if low > high {
		return NotExist
	}

	mid = low + (high-low)/2
	fmt.Println(mid)
	if l[mid] == item {
		return mid
	}
	if l[mid] < item {
		return F(l[mid+1:], item)
	} else {
		return F(l[:mid-1], item)
	}
}
