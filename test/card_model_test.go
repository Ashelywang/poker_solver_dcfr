package test

import (
	"github.com/Ashelywang/poker_solver_dcfr/src/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardModel(t *testing.T) {
	suits := []string{"s", "h", "c", "d"}
	strRanks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	for i, suit := range suits {
		for j, rank := range strRanks {
			card := common.NewCard(rank + suit)
			assert.Equal(t, card, common.Card(i * 13 + j))
		}
	}
}

func TestCardModelToInt32(t *testing.T) {
	suits := []string{"s", "h", "c", "d"}
	strRanks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	for i, suit := range suits {
		for j, rank := range strRanks {
			card := common.NewCard(rank + suit)
			assert.Equal(t, card.ToInt32(), int32(i * 13 + j))
		}
	}
}