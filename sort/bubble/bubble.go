/*
 * Revision History:
 *     Initial: 2018/05/10        Jon Snow
 */

package main

import (
	"math/rand"
	"fmt"
)

var count int
var opttimes int
func main() {
	var array []int
	for i := 0; i < 500; i++ {
		array = append(array, rand.Intn(1000))
	}
	//fmt.Println("排序前:", array)
	BubbleSort(array)
	fmt.Println("排序后:", array )
	fmt.Println("比较次数:", count,"\n操作次数:",opttimes)
}

func BubbleSort(array []int) {
	var flag bool
	for i := len(array); i >= 0; i-- {
		flag = true
		for j := 0; j < i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = false
				opttimes++
			}
			count++
		}
		if flag { break }
	}
}
