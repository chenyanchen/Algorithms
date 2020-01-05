/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var count int = 0

func main() {
	//slice := []int{}
	//for i := 0; i < 1000; i++ {
	//	slice = append(slice, rand.Intn(10000))
	//}
	//t := time.Now()
	////qsort(slice)
	//quickSort(slice, 0, len(slice)-1)
	//fmt.Printf("快速排序耗时:%v	比较次数:%v\n", time.Since(t), count)
	//fmt.Println("排序后:", slice)

	var array []int
	for i := 0; i < 500; i++ {
		array = append(array, rand.Intn(1000))
	}
	fmt.Println("排序前:", array)
	quickSort(array, 0, len(array)-1)
	fmt.Println("排序后:", array)
	fmt.Println("比较次数:", count)
}

func Quick_Sort(slice []int) {
	t := time.Now()
	//qsort(slice)
	quickSort(slice, 0, len(slice)-1)
	fmt.Printf("快速排序耗时:%v	比较次数:%v\n", time.Since(t), count)
}

// 优化版: 交换左边比中大右边比中小的元素
func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
				count++
			}
			for arr[j] > key {
				j--
				count++
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
			count++
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
			count++
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
				count++
				break
			}
			high--
		}

		for low < high {
			if key < s[low] {
				s[high] = s[low]
				count++
				break
			}
			low++
		}
	}
	s[low] = key

	quick1(s[:low-1])
	quick1(s[high+1:])
}
