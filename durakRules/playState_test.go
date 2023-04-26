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

func TestDefend(t *testing.T) {
	ps := PlayState{
		HandA: MakeHandFromCards(getCards("AS KH TD 2H 2C 4C")),
		HandB: MakeHandFromCards(getCards("KD TS 3H 4D 2S 3C")),
	}

	_, err := ps.Defend(MakeHandFromCards(getCards("2C")))
	if err == nil || err.ErrorType != IncorrectTurnForDefend {
		t.Error("Expecting IncorrectTurnForDefend, have ", err)

	}

	ps, erra := ps.Attack(MakeHandFromCards(getCards("2C")))
	if erra != nil {
		t.Error("valid attack returned error ", erra)
	}

	_, err = ps.Defend(MakeHandFromCards(getCards("3C")))
	if err != nil {
		t.Error("vallid defence returned error ", err)
	}

	_, err = ps.Defend(Hand(0))
	if err == nil || err.ErrorType != NotExactlyOneDefedCard {
		t.Error("Expecting NotExactlyOneDefendCard, have ", err)
	}

	_, err = ps.Defend(MakeHandFromCards(getCards("3C")))
	if err != nil {
		t.Error("vallid defence returned error ", err)
	}

	_, err = ps.Defend(MakeHandFromCards(getCards("7S")))
	if err == nil || err.ErrorType != HandBNotHoldingCard {
		t.Error("Expecting HandBNotHoldingCard, have ", err)
	}
}

func TestPull(t *testing.T) {
	ps := PlayState{
		HandA: MakeHandFromCards(getCards("AS KH TD 2H 2C 4C")),
		HandB: MakeHandFromCards(getCards("KD TS 3H 4D 2S 3C")),
	}

	_, _, err := ps.Pull()
	if err == nil || err.ErrorType != IncorrectTurnForPull {
		t.Error("Expecting IncorrectTurnForPull, have ", err)

	}

	ps, erra := ps.Attack(MakeHandFromCards(getCards("2C")))
	if erra != nil {
		t.Error("valid attack returned error ", erra)
	}

	_, _, err = ps.Pull()
	if erra != nil {
		t.Error("valid pull returned error ", erra)
	}
}

func TestYield(t *testing.T) {
	ps := PlayState{
		HandA: MakeHandFromCards(getCards("AS KH TD 2H 2C 4C")),
		HandB: MakeHandFromCards(getCards("KD TS 3H 4D 2S 3C")),
	}

	_, _, err := ps.Yield()
	if err == nil || err.ErrorType != IncorrectTurnForYield {
		t.Error("Expecting IncorrectTurnForYield, have ", err)
	}

	ps, erra := ps.Attack(MakeHandFromCards(getCards("2C")))
	if erra != nil {
		t.Error("valid attack returned error ", erra)
	}

	_, _, err = ps.Yield()
	if err == nil || err.ErrorType != IncorrectTurnForYield {
		t.Error("Expecting IncorrectTurnForYield, have ", err)
	}
	ps, errD := ps.Defend(MakeHandFromCards(getCards("3C")))
	if errD != nil {
		t.Error("vallid defence returned error ", errD)
	}

	_, _, err = ps.Yield()
	if err != nil {
		t.Error("vallid yield returned error ", err)
	}

}
