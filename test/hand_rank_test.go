package test

import (
	"github.com/Ashelywang/poker_solver_dcfr/src/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertHandsToKey(t *testing.T) {
	hand := []common.Card{common.NewCard("As"), common.NewCard("Ks"), common.NewCard("Qs"), common.NewCard("Js"), common.NewCard("Ts")}
	key := common.ConvertHandsToKey(hand)
	assert.Equal(t, key, string("As-Js-Ks-Qs-Ts"))

	hand = []common.Card{common.NewCard("2s"), common.NewCard("4c"), common.NewCard("Qs"), common.NewCard("Qh"), common.NewCard("Ts")}
	key = common.ConvertHandsToKey(hand)
	assert.Equal(t, key, string("2s-4c-Qh-Qs-Ts"))

	hand = []common.Card{common.NewCard("6s"), common.NewCard("6c"), common.NewCard("6h"), common.NewCard("Qh"), common.NewCard("Qs")}
	key = common.ConvertHandsToKey(hand)
	assert.Equal(t, key, string("6c-6h-6s-Qh-Qs"))
}

func TestCompareHands(t *testing.T) {
	hand1 := []common.Card{common.NewCard("8c"), common.NewCard("9s")}
	hand2 := []common.Card{common.NewCard("Ah"), common.NewCard("Qs")}
	board := []common.Card{common.NewCard("Tc"), common.NewCard("Jc"), common.NewCard("Kd"), common.NewCard("3h"), common.NewCard("7s")}
	hr := common.NewHandRanker("/Users/admin/ashe_workspace/poker_solver_dcfr/data/card5_dic_sorted.txt")
	rank1, _ := hr.GetHandRank(hand1, board)
	assert.Equal(t, rank1, int64(1603))
	rank2, _ := hr.GetHandRank(hand2, board)
	assert.Equal(t, rank2, int64(1600))
	result, _ := hr.CompareHands(hand1, hand2, board)
	assert.Equal(t, result, common.CompareResultLose)
}