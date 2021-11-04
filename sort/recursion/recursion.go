// Copyright 2020 Singularity, Inc. All rights reserved.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(F1(7))
	fmt.Println(F2(6, 1))

	fmt.Println(fib1(7))
	fmt.Println(fib2(7, 1, 1))
}

// 线性递归
func F1(n uint) uint {
	if n < 2 {
		return 1
	}
	return n * F1(n-1)
}

// 尾递归
func F2(n, s uint) uint {
	if n == 1 {
		return s
	}
	return F2(n-1, n*s)
}

// fib 线性递归
func fib1(n int) int {
	if n < 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// fibonacci 尾递归
func fib2(n, r1, r2 uint) uint {
	if n < 2 {
		return r1
	} else {
		return fib2(n-1, r2, r1+r2)
	}
}
