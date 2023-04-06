package test

import (
	"github.com/Ashelywang/poker_solver_dcfr/src/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardModel(t *testing.T) {
	suits := []string{"s", "h", "c", "d"}
	strRanks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
	for i, suit := range suits {
		for j, rank := range strRanks {
			card := common.NewCard(suit + rank)
			assert.Equal(t, card, common.Card(i * 13 + j + 1))
		}
	}
}