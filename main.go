package main

import (
	"fmt"
	"github.com/Ashelywang/poker_solver_dcfr/src/common"
)

func main() {
	suits := []string{"c", "h", "d", "s"}
	strRanks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}
	for _, suit := range suits {
		for _, rank := range strRanks {
			card := common.NewCard(suit + rank)
			fmt.Printf("Card: %s%s Num: %d \n", suit, rank, card)
		}
	}
}
