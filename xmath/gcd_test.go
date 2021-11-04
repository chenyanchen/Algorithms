// Copyright 2020 Singularity, Inc. All rights reserved.

package xmath_test

import (
	"testing"

	"github.com/chenyanchen/Algorithms/xmath"
)

func TestGCD(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				x: 1,
				y: 3,
			},
			want: 1,
		}, {
			name: "case 2",
			args: args{
				x: 6,
				y: 8,
			},
			want: 2,
		}, {
			name: "case 3",
			args: args{
				x: 1 << 10,
				y: 1 << 11,
			},
			want: 1 << 10,
		}, {
			name: "case 4",
			args: args{
				x: 1234567890,
				y: 1234567891,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := xmath.GCD(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("GCD() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
