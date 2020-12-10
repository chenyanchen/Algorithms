// Copyright 2020 Singularity, Inc. All rights reserved.

package list

import (
	"reflect"
	"testing"
)

func buildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var (
		root = &ListNode{Val: arr[0]}
		pre  *ListNode
	)
	pre = root
	for i := 1; i < len(arr); i++ {
		cur := ListNode{Val: arr[i]}
		pre.Next = &cur
		pre = pre.Next
	}
	return root
}

func pickUpList(root *ListNode) []int {
	arr := make([]int, 0)
	cur := root
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "normal case 1",
			args: args{buildList([]int{1, 2, 3, 4, 5})},
			want: buildList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", pickUpList(got), pickUpList(tt.want))
			}
		})
	}
}
