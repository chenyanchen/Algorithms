// Copyright 2020 Singularity, Inc. All rights reserved.

// Revision History:
//     Initial: 2019-05-15 14:02    Jon Snow
//     Description: 纸牌堆叠游戏, 手上有一堆纸牌,
//     一张放桌上, 一张放牌堆底部, 最后所有纸牌都在桌上,
//     请恢复原纸牌顺序
//     Example: 手上有一堆牌[1, 2, 3, 4, 5, 6, 7, 8],
//     按如上规则执行后, 桌面上的顺序为[1, 3, 5, 7, 2, 6, 4, 8]
//     请恢复手牌原顺序
package cardgames

type PileGame struct {
	Cards []int
}

func (g *PileGame) Do() (result []int) {
	result = g.Cards
	for i := 0; i < len(result); {
		if i%2 == 0 {
			i++
		} else {
			a := result[i]
			result = append(result[:i], result[i+1:]...)
			result = append(result, a)
		}
	}
	return
}

func (g *PileGame) Recover() (result []int) {
	result = make([]int, 0, len(g.Cards))

	return
}
