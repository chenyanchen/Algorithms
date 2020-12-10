// Copyright 2020 Singularity, Inc. All rights reserved.

/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

package main

import (
	"fmt"
)

func main() {
	a := []int{34, 52, 12, 45, 56, 10, 35, 37, 40}
	fmt.Println("排序前:", a)
	HeapSort(a)
	fmt.Println("排序后:", a)
	print(a)
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 构建大顶堆
func adjustHeap(l []int, start int) {
	// 将非叶子节点与其左叶子比较
	if l[start] < l[start*2+1] {
		swap(l, start, start*2+1)
	}
	// 将非叶子节点与其右叶子比较
	if l[start] < l[start*2+2] {
		swap(l, start, start*2+2)
	}
}

func HeapSort(l []int) {
	length := len(l)
	if length < 2 {
		return
	}

	for j := length; j > 0; j-- {
		// 从最后一个非叶子节点开始构建大顶堆
		for i := (j-1)/2 - 1; i >= 0; i-- {
			adjustHeap(l, i)
		}
		// 将根与最后一个元素交换
		swap(l, 0, j-1)
	}

}

func print(s []int) {
	fmt.Printf("			%d\n", s[0])
	fmt.Printf("	%v				%v\n", s[1], s[2])
	fmt.Printf("%v		%v		%v		%v\n", s[3], s[4], s[5], s[6])
}
