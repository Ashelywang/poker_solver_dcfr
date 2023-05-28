package common

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand []Card

type Board []Card

type HandRanker struct {
	Cards2Rank map[string]int64
}

type CompareResult int8

const (
	CompareResultUnknown CompareResult = 0
	CompareResultWin CompareResult = 1
	CompareResultLose CompareResult = 2
	CompareResultTie  CompareResult = 3
)

func NewHandRanker(filePath string) *HandRanker {
	// Load hand ranker from filePath
	ranker := HandRanker{Cards2Rank: map[string]int64{}}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return nil
	}
	defer file.Close()

	// Read each line from file

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fmt.Printf("Line has more than 2 parts: %s\n", line)
			continue // Skip lines that don't have a key-value format
		}
		key := parts[0]
		value := parts[1]
		num, err2 := strconv.ParseInt(value, 10, 64)
		if err2 == nil {
			ranker.Cards2Rank[key] = num
		} else {
			fmt.Printf("Failed to parse value: %s, err: %v\n", value, err2)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while scanning file: %s\n", err)
		return nil
	}
	fmt.Printf("Finished scanning file, size: %d\n", len(ranker.Cards2Rank))
	return &ranker
}

func (hr *HandRanker) CompareHands(h1 Hand, h2 Hand, board Board) (CompareResult, error) {
	if len(h1) != 2 || len(h2) != 2 {
		return CompareResultUnknown, fmt.Errorf("hands must have 2 cards")
	}
	if len(board) != 5 {
		return CompareResultUnknown, fmt.Errorf("board must have 5 cards")
	}

	handRank1, err := hr.GetHandRank(h1, board)
	if err != nil || handRank1 == 0 {
		return CompareResultUnknown, err
	}
	handRank2, err := hr.GetHandRank(h2, board)
	if err != nil || handRank2 == 0 {
		return CompareResultUnknown, err
	}
	if handRank1 < handRank2 {
		return CompareResultWin, nil
	}
	if handRank1 > handRank2 {
		return CompareResultLose, nil
	}
	if handRank1 == handRank2 {
		return CompareResultTie, nil
	}
	return CompareResultUnknown, nil
}

func (hr *HandRanker) GetHandRank (hand Hand, board Board) (int64, error) {
	if len(hand) != 2 {
		return 0, fmt.Errorf("hand must have 2 cards")
	}
	if len(board) != 5 {
		return 0, fmt.Errorf("board must have 5 cards")
	}

	//	Generate all possible 5 cards combination to get all 21 possible combinations
	allCards := append(hand, board...)
    possibleCombinations := GenerateSubset5(allCards)

	//	Find the best hand
	bestRank := int64(7463)
	for _, combination := range possibleCombinations {
		//	Find the best hand in this combination
		key := ConvertHandsToKey(combination)
		if key == "" {
			continue
		}
		rank, ok := hr.Cards2Rank[key]
		if !ok {
			continue
		}
		if rank < bestRank {
			bestRank = rank
		}
	}
	if bestRank == 7463 {
		return 0, fmt.Errorf("failed to find best rank")
	}
	return bestRank, nil
}

func GenerateSubset5(cards []Card) [][]Card {
	var result [][]Card
	for i := 0; i < len(cards); i++ {
		for j := i+1; j < len(cards); j++ {
			for k := j+1; k < len(cards); k++ {
				for l := k+1; l < len(cards); l++ {
					for m := l+1; m < len(cards); m++ {
						result = append(result, []Card{cards[i], cards[j], cards[k], cards[l], cards[m]})
					}
				}
			}
		}
	}
	return result
}

func ConvertHandsToKey(cards []Card) string {
	if len(cards) != 5 {
		return ""
	}
	key := []string{}
	for _, card := range cards {
		key =append(key, card.ToString())
	}
	sort.Strings(key)
	// Join the key with "-"
	result := strings.Join(key, "-")
	return result
}