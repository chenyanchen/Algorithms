// Copyright 2020 Singularity, Inc. All rights reserved.

// Reversion:
//   Init: Jon Snow    2020/4/16 21:54

package quicksort

import (
	"math/rand"
	"testing"
)

func TestQ(t *testing.T) {
	var tests = []struct {
		input []int
	}{
		{rand.Perm(10)},
	}
	for _, test := range tests {
		q(test.input, 0, len(test.input)-1)
		for i := range test.input {
			if i < len(test.input)-1 {
				if test.input[i] > test.input[i+1] {
					t.Fatalf("got: %v", test.input)
				}
			}
		}
	}
}
