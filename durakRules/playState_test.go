package durakRules

import (
	"strings"
	"testing"
)

func TestAttack(t *testing.T) {
	ps := PlayState{
		HandA: MakeHandFromCards(getCards("AS KH TD 2H 2C 4C")),
		HandB: MakeHandFromCards(getCards("KD TS 3H 4D 2S 8D")),
	}
	ps2, err := ps.Attack(MakeHandFromCards(getCards("2C")))
	if err != nil {
		t.Error("valid attack returned error ", err)
	}

	_, err = ps2.Attack(MakeHandFromCards(getCards("2C 4C")))
	if err == nil || err.ErrorType != IncorrectTurnForAttack {
		t.Error("Expecting IncorrectTurnForAttack, have ", err)
	}

	_, err = ps.Attack(MakeHandFromCards(getCards("2C 4C")))
	if err == nil || err.ErrorType != NotExactlyOneAttackCard {
		t.Error("Expecting NotExactlyOneAttackCard, have ", err)
	}

	_, err = ps.Attack(MakeHandFromCards(getCards("3S")))
	if err == nil || err.ErrorType != HandANotHoldingCard {
		t.Error("Expecting HandANotHoldingCard, have ", err)
	}
	ps2.DefendCards = MakeHandFromCards(getCards("2S"))

	_, err = ps2.Attack(MakeHandFromCards(getCards("4C")))
	if err == nil || err.ErrorType != CardNotOnTable {
		t.Error("Expecting CardNotOnTable, have ", err)
	}
	_, err = ps2.Attack(MakeHandFromCards(getCards("2H")))

	if err != nil {
		t.Error("valid attack returned error ", err)
	}
}

func getCards(s string) []Card {
	//parses a list of cards like 3H 2S KD AS into a slice of Card
	cs := strings.Split(s, " ")
	cards := make([]Card, len(cs))
	for i, c := range cs {
		cards[i].Rank = Rank(strings.Index("23456789TJQKA", string(c[0])))
		cards[i].Suit = Suit(strings.Index("HCDS", string(c[1])))
	}
	return cards
}