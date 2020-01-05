/*
 * Revision History:
 *     Initial: 2018/05/16        Chen Yanchen
 */

package main

import (
	"fmt"
	"math/rand"
)

var count int

func main() {
	var array []int
	for i := 0; i < 500; i++ {
		array = append(array, rand.Intn(1000))
	}
	//fmt.Println("排序前:", array)
	ShellSort(array)
	//fmt.Println("排序后:", array )
	fmt.Println("比较次数:", count)
}

func ShellSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}
	var k int
	for k < length/3 { //寻找合适的间隔h
		k = 3*length + 1
	}
	// 按 k 步间隔排序
	for ; k >= 1; k /= 3 {
		for i := k; i < length; i++ {
			for j := i; j >= k; j -= k {
				if array[j] < array[j-k] {
					array[j], array[j-k] = array[j-k], array[j]
				}
				count++
			}
		}
	}
}

func shellSort(array []int) {
	n := len(array)
	if n < 2 {
		return
	}
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && array[j] < array[j-key] {
				array[j], array[j-key] = array[j-key], array[j]
				j = j - key
				count++
			}
		}
		key = key / 2
	}
}
