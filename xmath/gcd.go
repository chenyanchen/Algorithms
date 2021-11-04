// Copyright 2020 Singularity, Inc. All rights reserved.

package xmath

// GCD return the Greatest Common Divisor for 2 int.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
