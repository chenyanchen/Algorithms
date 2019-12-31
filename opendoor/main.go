// Revision History:
//     Initial: 2019-06-19 11:11    Jon Snow
//     题目描述: 有3个门, 其中一扇门后有奖品, 现在你先选一扇门,
//     此时主持人打开一扇空门, 此时再给你一次机会选则剩下的两扇门,
//     问选择换一扇门还是不换获得奖品的概率大
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count, right := 100000, 0
	for i := 0; i < count; i++ {
		if simulate(false) {
			right++
		}
	}
	fmt.Println("no change rate:", float64(right*100)/float64(count))

	right = 0
	for i := 0; i < count; i++ {
		if simulate(true) {
			right++
		}
	}
	fmt.Println("change rate:", float64(right*100)/float64(count))
}

// 模拟一次开门活动, change 为是否改变选择结果, 返回中奖结果
func simulate(change bool) bool {
	pool := make([]int, 3)
	rand.Seed(time.Now().UnixNano())
	pool[rand.Intn(3)] = 1

	choose := rand.Intn(3)

	out := 0
	for i, v := range pool {
		if i != choose && v != 1 { // 打开一扇空门
			out = i
			break
		}
	}

	if change {
		for i := range pool {
			if choose != i && out != i {
				choose = i
			}
		}
	}

	return pool[choose] == 1
}
