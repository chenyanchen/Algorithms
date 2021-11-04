// Copyright 2020 Singularity, Inc. All rights reserved.

package xmath

// SimplestRatio return simplest ratio for 2 int.
func SimplestRatio(a, b int) (int, int) {
	gcd := GCD(a, b)

	return a / gcd, b / gcd
}
