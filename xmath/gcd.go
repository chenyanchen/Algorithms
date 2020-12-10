// Copyright 2020 Singularity, Inc. All rights reserved.

package xmath

// Deprecated: without recursive
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

// GCD return the Greatest Common Divisor for 2 int.
func GCD(x, y int) int {
	mod := x % y
	for mod > 0 {
		x, y = y, mod
		mod = x % y
	}
	return y
}
