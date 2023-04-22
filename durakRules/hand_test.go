package durakRules

import (
	"testing"
)

func TestGetCards(t *testing.T) {
	h1 := Hand(1 << (int(Clubs)*13 + int(Four)))
	cs1 := h1.GetCards()
	if len(cs1) != 1 {
		t.Error("Not just one card but ", len(cs1), " cards")
	}
	if cs1[0].Suit != Clubs {
		t.Error("Card not clubs ", cs1)
	}
	if cs1[0].Rank != Four {
		t.Error("Card not four ", cs1)
	}

	h2 := MakeHandFromCards([]Card{{Two, Hearts}, {Ace, Clubs}})
	cs2 := h2.GetCards()
	if len(cs2) != 2 {
		t.Error("Not exactly two cards but ", len(cs2), " cards")
	}
	if cs2[1].Suit != Clubs {
		t.Error("Card not clubs ", cs2)
	}
	if cs2[1].Rank != Ace {
		t.Error("Card not ace", cs2)
	}
}

func TestCombine(t *testing.T) {
	h1 := MakeHandFromCards([]Card{{Two, Hearts}, {Three, Clubs}, {King, Clubs}})
	h2 := MakeHandFromCards([]Card{{Two, Spades}})
	h3 := MakeHandFromCards([]Card{{Jack, Spades}})
	h4 := h1.Combine(h2).Combine(h3).GetCards()
	if len(h4) != 5 {
		t.Error("error combining, see tests code")
	}

}

func TestReduceToRank(t *testing.T) {
	h1 := MakeHandFromCards([]Card{{Two, Hearts}, {Three, Clubs}, {King, Clubs}})
	h2 := MakeHandFromCards([]Card{{Two, Spades}})
	h3 := MakeHandFromCards([]Card{{Jack, Spades}})
	h4 := h1.Combine(h2).Combine(h3).reduceToRank().GetCards()
	if len(h4) != 4 {
		t.Error("error combining and reducing, see tests code")
	}
}

func TestHasRanks(t *testing.T) {
	h1 := MakeHandFromCards([]Card{{Two, Hearts}, {Three, Clubs}, {King, Clubs}})
	h2 := MakeHandFromCards([]Card{{Two, Spades}})
	h3 := MakeHandFromCards([]Card{{Jack, Spades}})

	if !h1.HasRanks(h2) {
		t.Error("aparently ", h1.GetCards(), " doesn't have rank", h2.GetCards())
	}
	if h1.HasRanks(h3) {
		t.Error("aparently ", h1.GetCards(), " has rank", h3.GetCards())
	}

}