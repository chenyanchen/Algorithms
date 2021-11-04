// Copyright 2020 Singularity, Inc. All rights reserved.

// Reversion:
//   Init: Jon Snow    2020/4/16 21:54

package quicksort_test

import (
	"math/rand"
	"testing"

	"github.com/chenyanchen/Algorithms/sort/quicksort"
)

func TestQ(t *testing.T) {
	var tests = []struct {
		input []int
	}{
		{rand.Perm(10)},
	}
	for _, test := range tests {
		quicksort.Sort(test.input, 0, len(test.input)-1)
		for i := range test.input {
			if i < len(test.input)-1 {
				if test.input[i] > test.input[i+1] {
					t.Fatalf("got: %v", test.input)
				}
			}
		}
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				arr:   []int{3, 3, 3, 3, 2, 1, 1, 1},
				start: 0,
				end:   0,
			},
			want: []int{1, 1, 1, 2, 3, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
			},
		)
	}
}
