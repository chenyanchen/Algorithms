// Copyright 2020 Singularity, Inc. All rights reserved.

package xmath_test

import (
	"testing"

	"github.com/chenyanchen/Algorithms/xmath"
)

func TestSimplestRatio(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "case 1",
			args: args{
				a: 1,
				b: 1,
			},
			want:  1,
			want1: 1,
		}, {
			name: "case 2",
			args: args{
				a: 99,
				b: 99,
			},
			want:  1,
			want1: 1,
		}, {
			name: "case 3",
			args: args{
				a: 2,
				b: 4,
			},
			want:  1,
			want1: 2,
		}, {
			name: "case 4",
			args: args{
				a: 4,
				b: 8,
			},
			want:  1,
			want1: 2,
		}, {
			name: "case 5",
			args: args{
				a: 6,
				b: 8,
			},
			want:  3,
			want1: 4,
		}, {
			name: "case 6",
			args: args{
				a: 1234567890,
				b: 1234567891,
			},
			want:  1234567890,
			want1: 1234567891,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, got1 := xmath.SimplestRatio(tt.args.a, tt.args.b)
				if got != tt.want {
					t.Errorf("SimplestRatio() got = %v, want %v", got, tt.want)
				}
				if got1 != tt.want1 {
					t.Errorf("SimplestRatio() got1 = %v, want %v", got1, tt.want1)
				}
			},
		)
	}
}
