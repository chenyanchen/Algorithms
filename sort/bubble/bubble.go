/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

 package main

import (
	"fmt"
)

func main() {
	a := []int{34, 5, 3, 6, 9, 0, 12, 23, 13, 7}
	fmt.Println("排序前:", a)
	BubbleSort(a)
	fmt.Println("排序后:", a)
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func BubbleSort(s []int) {
	for i := len(s); i >= 0; i-- {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j, j+1)
			}
		}
	}
}
