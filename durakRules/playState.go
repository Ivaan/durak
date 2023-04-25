package durakRules

import "fmt"

type PlayState struct {
	HandA       Hand
	HandB       Hand
	AttackCards Hand
	DefendCards Hand
	Discarded   Hand
}

const FirstSuitStartBit = 0
const SecondSuitStartBit = 13
const ThirdSuitStartBit = 26
const TrumpSuitStartBit = 39

func (ps PlayState) Attack(p Hand) (PlayState, *AttackError) {
	if ps.AttackCards.CardCount() != ps.DefendCards.CardCount() {
		return ps, &AttackError{ErrorType: IncorrectTurnForAttack, PlayState: &ps, Play: &p}
	}
	if p.CardCount() != 1 {
		return ps, &AttackError{ErrorType: NotExactlyOneAttackCard, PlayState: &ps, Play: &p}
	}
	if !ps.HandA.HasCards(p) {
		return ps, &AttackError{ErrorType: HandANotHoldingCard, PlayState: &ps, Play: &p}
	}
	if ps.AttackCards.CardCount() > 0 && !ps.AttackCards.HasRanks(p) {
		return ps, &AttackError{ErrorType: CardNotOnTable, PlayState: &ps, Play: &p}
	}

	ps.HandA ^= p       //remove played card from Hand A
	ps.AttackCards |= p //add played card to attack cards

	return ps, nil
}

func (ps PlayState) Defend(p Hand) (PlayState, *DefendError) {
	if ps.AttackCards.CardCount() != ps.DefendCards.CardCount()+1 {
		return ps, &DefendError{ErrorType: IncorrectTurnForDefend, PlayState: &ps}
	}
	if p.CardCount() != 1 {
		return ps, &DefendError{ErrorType: NotExactlyOneDefedCard, PlayState: &ps, Play: &p}
	}
	if !ps.HandB.HasCards(p) {
		return ps, &DefendError{ErrorType: HandBNotHoldingCard, PlayState: &ps, Play: &p}
	}

	ps.HandB ^= p
	newDefendCards := ps.DefendCards | p

	attackCards := ps.AttackCards.GetCards()
	defendCards := newDefendCards.GetCards()
	ai := 0
	di := 0
	cardsToBeTrumped := 0
	for {
		ac := attackCards[ai]
		dc := defendCards[di]
		if ac.Suit < dc.Suit {
			cardsToBeTrumped++
			ai++
		} else if ac.Suit > dc.Suit {
			return ps, &DefendError{ErrorType: NotAllCardsBeaten, PlayState: &ps, Play: &p}
		} else { //same suit
			if ac.Rank < dc.Rank {
				ai++
				di++
			} else {
				return ps, &DefendError{ErrorType: NotAllCardsBeaten, PlayState: &ps, Play: &p}
			}
		}
		if ai == len(attackCards) {
			if di == len(defendCards) && cardsToBeTrumped == 0 {
				break
			}
			if cardsToBeTrumped == len(defendCards)-di && dc.Suit == Trump {
				break
			}
			return ps, &DefendError{ErrorType: NotAllCardsBeaten, PlayState: &ps, Play: &p}
		}
	}

	ps.DefendCards = newDefendCards
	return ps, nil
}

func (ps PlayState) Pull() (PlayState, DrawDirective, *PullError) {
	dd := DrawDirective{}

	if ps.AttackCards.CardCount() != ps.DefendCards.CardCount()+1 {
		return ps, dd, &PullError{ErrorType: IncorrectTurnForPull, PlayState: &ps}
	}
	ps.HandB = ps.HandB.Combine(ps.AttackCards.Combine(ps.DefendCards))

	dd.NumberOfCardsToDrawIntoHandA = computeDraw(ps.HandA.CardCount())
	dd.NumberOfCardsToDrawIntoHandB = computeDraw(ps.HandB.CardCount())

	return ps, dd, nil
}

func (ps PlayState) Yield() (PlayState, DrawDirective, *YieldError) {
	dd := DrawDirective{}

	if ps.AttackCards.CardCount() != ps.DefendCards.CardCount() {
		return ps, dd, &YieldError{ErrorType: IncorrectTurnForYield, PlayState: &ps}
	}
	ps.Discarded = ps.Discarded.Combine(ps.AttackCards.Combine(ps.DefendCards))
	ps.AttackCards = Hand(0)
	ps.DefendCards = Hand(0)
	dd.NumberOfCardsToDrawIntoHandA = computeDraw(ps.HandA.CardCount())
	dd.NumberOfCardsToDrawIntoHandB = computeDraw(ps.HandB.CardCount())

	return ps, dd, nil

}

func (ps PlayState) String() string {
	return fmt.Sprintf(
		"HandA:%s\nHandB:%s\nAttack:%s\nDefend:%s\nDiscarded:%s\n",
		ps.HandA.GetCards(),
		ps.HandB.GetCards(),
		ps.AttackCards.GetCards(),
		ps.DefendCards.GetCards(),
		ps.Discarded.GetCards(),
	)
}

func computeDraw(c int) int {
	if c < 6 {
		return 6 - c
	} else {
		return 0
	}
}

var eightBitBitCounts [256]int

func genEightBitBitCounts() [256]int {
	var o [256]int
	for i := 0; i < 256; i++ {
		v := 0
		for b := 0; b < 8; b++ {
			if (i>>b)&1 == 1 {
				v++
			}
		}
		o[i] = v
	}
	return o
}
