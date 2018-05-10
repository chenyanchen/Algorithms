/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

 package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 8, 5, 9, 4, 3, 6, 6, 7, 14, 36, -1}
	fmt.Println("排序前:", slice)
	QuickSort(slice)
	fmt.Println("排序后:", slice)
}

func Quick(s []int) {
	var small, big []int

	if len(s) < 2 {
		return
	} else {
		pivot := s[0]
		for _, v := range s[1:] {
			if pivot > v {
				small = append(small, v)
				fmt.Println(small)
			} else {
				big = append(big, v)
			}
		}
		Quick(small)
		Quick(big)
		s = append(append(small, pivot), big...)
	}
}


// Quick1 将选中的元素填充至初始位置
func Quick1(s []int) {
	var key, low, high int
	key = s[0]
	low = 0
	high = len(s) - 1

	if len(s) < 2 {
		return
	}

	for low < high {

		for low < high {
			if key > s[high] {
				s[low] = s[high]
				break
			}
			high--
		}

		for low < high {
			if key < s[low] {
				s[high] = s[low]
				break
			}
			low++
		}
	}
	s[low] = key

	QuickSort(s[:low-1])
	QuickSort(s[high+1:])
}

// 交换选中的元素
func QuickSort(s []int) {
	var key, low, high int
	key = s[0]
	low = 0
	high = len(s) - 1

	if len(s) < 2 {
		return
	}

	for low < high {

		for low < high {
			if key > s[high] {
				swap(s, low, high)
				break
			}
			high--
		}

		for low < high {
			if key < s[low] {
				swap(s, low, high)
				break
			}
			low++
		}
	}
	s[low] = key

	QuickSort(s[:low-1])
	QuickSort(s[high+1:])
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}
