package durakRules

import (
	"encoding/binary"
	"fmt"
)

type Hand uint64

type Suit int

const (
	Hearts Suit = iota
	Clubs
	Diamonds
	Spades
	Trump = Spades
)

func (s Suit) String() string {

	switch s {
	case Hearts:
		return "Hearts"
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Spades:
		return "Spades"
	default:
		panic("unknown Suit" + fmt.Sprintf("%v", int(s)))
	}
}

type Rank int

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (r Rank) String() string {
	switch r {

	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		panic("unknown Rank" + fmt.Sprintf("%v", int(r)))
	}
}

type Card struct {
	Rank
	Suit
}

func MakeHandFromCards(cards []Card) Hand {
	h := uint64(0)
	for _, c := range cards {
		h |= 1 << (int(c.Suit)*13 + int(c.Rank))
	}
	return Hand(h)
}

func (h Hand) GetCards() []Card {
	cards := make([]Card, 0)
	for s := Hearts; s <= Spades; s++ {
		for r := Two; r <= Ace; r++ {
			if (h>>(int(s)*13+int(r)))&1 == 1 {
				cards = append(cards, Card{Rank: r, Suit: s})
			}
		}
	}
	return cards
}

func (h Hand) HasCards(h2 Hand) bool {
	return (h & h2) == h2
}

func (h Hand) HasRanks(h2 Hand) bool {
	hRanks := h.reduceToRank()
	h2Ranks := h2.reduceToRank()

	return (hRanks & h2Ranks) == h2Ranks
}

func (h Hand) Combine(h2 Hand) Hand {
	return h | h2
}

func (h Hand) CardCount() int {
	c := 0
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(h))
	for _, b := range bs {
		c += eightBitBitCounts[b]
	}
	return c
}

type DrawDirective struct {
	NumberOfCardsToDrawIntoHandA int
	NumberOfCardsToDrawIntoHandB int
}

func (h Hand) reduceToRank() Hand {
	o := uint64(h) & 0x0000000000001FFF //1FFF is 13 bits
	o |= uint64(h) >> 13 & 0x0000000000001FFF
	o |= uint64(h) >> 26 & 0x0000000000001FFF
	o |= uint64(h) >> 39 & 0x0000000000001FFF

	return Hand(o)
}