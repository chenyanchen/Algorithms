// Copyright 2020 Singularity, Inc. All rights reserved.

package merge

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestMergeSort(t *testing.T) {
	num := rand.Intn(1000)
	arr := make([]int, num)
	for i := range arr {
		arr[i] = rand.Int()
	}

	result := MergeSort(arr)
	sort.Ints(arr)

	for i := range arr {
		if arr[i] != result[i] {
			t.Logf("expect: %d got: %d", arr[i], result[i])
		}
	}
}
