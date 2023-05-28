package common

type Card int32

var (
	intRanks [14]uint8
	strRanks = "23456789TJQKA"

	charRankToIntRank = map[uint8]uint8{}
	charSuitToInt = map[uint8]uint8{'s': 0, 'h':1, 'c':2, 'd':3}
	charIntToSuit = map[uint8]string{0: "s", 1:"h", 2:"c", 3:"d"}
	charIntToRank = map[uint8]string{0: "2", 1:"3", 2:"4", 3:"5", 4:"6", 5:"7", 6:"8", 7:"9", 8:"T", 9:"J", 10:"Q", 11:"K", 12:"A"}
)

func init() {
	for i := 1; i <= 13; i++ {
		charRankToIntRank[strRanks[i-1]] = uint8(i-1)
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
		return 255
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

	suit := SuitToNum(s[1])
	rank := OrderToNum(s[0])

	if suit == 255 || rank == 255 {
		return -1
	}
	return Card(suit*13 + rank)
}

func (c Card) ToString() string{
	rank := c % 13
	suit := c / 13
	if suit <= 3 && suit >=0 {
		result := charIntToRank[uint8(rank)] + charIntToSuit[uint8(suit)]
		return result
	}
	return ""
}

func (c Card) ToInt32() int32 {
	return int32(c)
}