package common

type Card int32

var (
	intRanks [14]uint8
	strRanks = "A23456789TJQK"

	charRankToIntRank = map[uint8]uint8{}
	charSuitToInt = map[uint8]uint8{'s': 0, 'h':1, 'c':2, 'd':3}
)

func init() {
	for i := 1; i <= 13; i++ {
		charRankToIntRank[strRanks[i-1]] = uint8(i)
	}
}

func SuitToNum(s uint8) uint8 {
	val, ok :=  charSuitToInt[s]
	if !ok {
		return 255
	}
	return val
}

func OrderToNum(s uint8) uint8 {
	val, ok :=  charRankToIntRank[s]
	if !ok {
		return 0
	}
	return val
}


func NewCard(s string) Card {
	// Club c
	// Spade s
	// Heart h
	// Diamond d
	if len(s) <= 1 {
		return -1
	}

	suit := SuitToNum(s[0])
	rank := OrderToNum(s[1])

	if suit == 255 || rank == 0 {
		return -1
	}
	return Card(suit*13 + rank)
}

