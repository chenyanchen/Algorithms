// Copyright 2020 Singularity, Inc. All rights reserved.

package main

import (
	"fmt"
	"math/rand"
)

var count int
var opttimes int

func main() {
	var array []int
	for i := 0; i < 500; i++ {
		array = append(array, rand.Intn(1000))
	}
	//fmt.Println("排序前:", array)
	SelectSort(array)
	//fmt.Println("排序后:", array)
	fmt.Println("比较次数:", count, "\n操作次数:", opttimes)
}

func SelectSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		key := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[key] {
				key = j
				opttimes++
			}
			count++
		}
		s[i], s[key] = s[key], s[i]
	}
}
