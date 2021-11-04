// Copyright 2020 Singularity, Inc. All rights reserved.

/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

package quicksort

// 优化版: 交换左边比中大右边比中小的元素
func Sort(arr []int, start, end int) {
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
			Sort(arr, start, j)
		}
		if end > i {
			Sort(arr, i, end)
		}
	}
}

// SortV2 from:https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F#Go
func SortV2(data []int) {
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
	SortV2(data[:head])
	SortV2(data[head+1:])
}

// 递归
func SortV3(s []int) []int {
	var small, big []int

	if len(s) < 2 {
		return s
	}

	pivot := s[0]
	for _, v := range s[1:] {
		if pivot > v {
			small = append(small, v)
		} else {
			big = append(big, v)
		}
	}
	SortV3(small)
	SortV3(big)
	s = append(append(small, pivot), big...)

	return s
}

// 最容易理解
func SortV4(s []int) {
	var key, left, right int

	if len(s) < 2 {
		return
	}

	key = s[0]
	left, right = 0, len(s)-1

	for left < right {
		for left < right {
			if key > s[right] {
				s[left] = s[right]

				break
			}
			right--
		}

		for left < right {
			if key < s[left] {
				s[right] = s[left]

				break
			}
			left++
		}
	}
	s[left] = key

	SortV4(s[:left-1])
	SortV4(s[right+1:])
}

func SortV5(arr []int, start, end int) {
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
		SortV5(arr, start, r)
	}
	if l < end {
		SortV5(arr, l, end)
	}
}
