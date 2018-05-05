package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 8, 5, 9, 4, 3, 6, 6, 7, 14, 36, -1}
	fmt.Println("排序前:", slice)
	SelectSort(slice)
	fmt.Println("排序后:", slice)
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SelectSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		key := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[key] {
				key = j
			}
		}
		swap(s, i, key)
	}
}
