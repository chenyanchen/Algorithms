package main

import (
	"fmt"
)

func main() {
	a := []int{34,5, 3, 6, 9, 0, 12, 23, 13, 7}
	fmt.Println("排序前:", a)
	InsertSort(a)
	fmt.Println("排序后:", a)
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func InsertSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		for j := i+1; j < len(s); j++ {
			if s[i] > s[j] {
				swap(s, i, j)
			}
		}
	}
}
