// Copyright 2020 Singularity, Inc. All rights reserved.

package xmath

// Deprecated:
func SimplestRatioV1(a, b int64) (int64, int64) {
	for i := int64(2); i <= a && i <= b; {
		if a%i == 0 && b%i == 0 {
			a, b = a/i, b/i
		} else {
			// TODO: optimize this part
			i++
		}
	}
	return a, b
}

// SimplestRatio return simplest ratio for 2 int.
func SimplestRatio(a, b int) (int, int) {
	gcd := GCD(a, b)
	return a / gcd, b / gcd
}
