// Copyright 2020 Singularity, Inc. All rights reserved.

// Reversion:
//   Init: Jon Snow    2020/4/16 22:09
//   Desc: 合并区间
//   Url: https://leetcode-cn.com/problems/merge-intervals/

package merge

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, 1)
	a, b := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if b >= intervals[i][0] {
			if b < intervals[i][1] {
				b = intervals[i][1]
			}
		} else {
			res = append(res, []int{a, b})
			a = intervals[i][0]
			b = intervals[i][1]
		}
	}
	res = append(res, []int{a, b})
	return res
}
