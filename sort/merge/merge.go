// Copyright 2020 Singularity, Inc. All rights reserved.

/*
 * Revision History:
 *     Initial: 2018/06/11        Jon Snow
 */

package merge

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
		} else {
			list = append(list, right[j])
			j++
		}
	}
	list = append(list, left[i:]...)
	list = append(list, right[j:]...)
	return list
}
