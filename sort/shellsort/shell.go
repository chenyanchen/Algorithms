/*
 * Revision History:
 *     Initial: 2018/05/16        Chen Yanchen
 */

package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 8, 5, 9, 4, 3, 6, 6, 7, 14, 36, -1, 23}
	fmt.Println("排序前:", slice)
	ShellSort(slice)
	fmt.Println("排序后:", slice)
}

func ShellSort(l []int) {
	length := len(l)
	if length < 2 {
		return
	}
	// 按 k 步间隔排序
	for k := length / 2; k > 0; k /= 2 {
		for i := k; i < length; i++ {
			for j := i; j >= k; j -= k {
				if l[j] < l[j-k] {
					l[j], l[j-k] = l[j-k], l[j]
				}
			}
		}

	}
}
