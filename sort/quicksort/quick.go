// Copyright 2020 Singularity, Inc. All rights reserved.

/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

package quicksort

// 优化版: 交换左边比中大右边比中小的元素
func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

// qsort from:https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F#Go
func qsort(data []int) {
	length := len(data)
	if length < 2 {
		return
	}
	mid := data[length/2]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

// 递归
func quick2(s []int) []int {
	var small, big []int

	if len(s) < 2 {
		return s
	} else {
		pivot := s[0]
		for _, v := range s[1:] {
			if pivot > v {
				small = append(small, v)
			} else {
				big = append(big, v)
			}
		}
		quick2(small)
		quick2(big)
		s = append(append(small, pivot), big...)
	}
	return s
}

// 最容易理解
func quick1(s []int) {
	var key, low, high int

	if len(s) < 2 {
		return
	}

	key = s[0]
	low = 0
	high = len(s) - 1

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

	quick1(s[:low-1])
	quick1(s[high+1:])
}

func q(arr []int, start, end int) {
	if start >= end {
		return
	}
	l, r := start, end
	mid := arr[(l+r)/2]
	for l <= r {
		for arr[l] < mid {
			l++
		}
		for mid < arr[r] {
			r--
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	if start < r {
		q(arr, start, r)
	}
	if l < end {
		q(arr, l, end)
	}
}
