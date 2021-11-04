// Copyright 2020 Singularity, Inc. All rights reserved.

package list_test

import (
	"reflect"
	"testing"

	"github.com/chenyanchen/Algorithms/list"
)

func buildList(arr []int) *list.Node {
	if len(arr) == 0 {
		return nil
	}

	var (
		// root *list.Node
		root = &list.Node{Val: arr[0]}
		pre  *list.Node
	)
	pre = root
	for i := 1; i < len(arr); i++ {
		cur := list.Node{Val: arr[i]}
		pre.Next = &cur
		pre = pre.Next
	}
	return root
}

func pickUpList(root *list.Node) []int {
	arr := make([]int, 0)
	cur := root
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}

func Test_Reverse(t *testing.T) {
	t.Parallel()

	type args struct {
		head *list.Node
	}
	tests := []struct {
		name string
		args args
		want *list.Node
	}{
		{
			name: "normal case 1",
			args: args{buildList([]int{1, 2, 3, 4, 5})},
			want: buildList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				if got := list.Reverse(tt.args.head); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Reverse() = %v, want %v", pickUpList(got), pickUpList(tt.want))
				}
			},
		)
	}
}
