/*
 * Revision History:
 *     Initial: 2018/06/11        Jon Snow
 */

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JonSnow47/Algorithms/sort/quicksort"
)

var count int = 0

func main() {
	for i := 1; i <= 3; i++ {
		nums := i * 500000
		var array []int
		for i := 0; i < nums; i++ {
			array = append(array, rand.Int())
		}
		fmt.Println("元素个数:", nums)
		t := time.Now()
		_ = MergeSort(array)
		fmt.Printf("归并排序耗时:%v	比较次数:%v\n", time.Since(t), count)
		//fmt.Println("排序后:", list)

		quicksort.Quick_Sort(array)
	}
}

func MergeSort(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}
	key := length / 2

	left := MergeSort(array[:key])
	right := MergeSort(array[key:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var list []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			list = append(list, left[i])
			i++
			count++
		} else {
			list = append(list, right[j])
			j++
			count++
		}
	}
	list = append(list, left[i:]...)
	list = append(list, right[j:]...)
	return list
}
