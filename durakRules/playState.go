package durakRules

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

type Play Hand

func (ps PlayState) Attack(p Hand) (PlayState, *AttackError) {
	if ps.AttackCards.CardCount() != ps.DefendCards.CardCount() {
		return ps, &AttackError{ErrorType: IncorrectTurnForAttack}
	}
	if p.CardCount() != 1 {
		return ps, &AttackError{ErrorType: NotExactlyOneAttackCard, Hand: &p}
	}
	if !ps.HandA.HasCards(p) {
		return ps, &AttackError{ErrorType: HandANotHoldingCard, Hand: &p}
	}
	if ps.AttackCards.CardCount() > 0 && !ps.AttackCards.HasRanks(p) {
		return ps, &AttackError{ErrorType: CardNotOnTable, Hand: &p}
	}

	ps.HandA ^= p       //remove played card from Hand A
	ps.AttackCards |= p //add played card to attack cards

	return ps, nil
}

func (ps PlayState) Defend(p Hand) (PlayState, *DefendError) {

}

func (ps PlayState) Pull() (PlayState, DrawDirective, *PullError) {
	dd := DrawDirective{}

	if ps.AttackCards.CardCount() != ps.DefendCards.CardCount()+1 {
		return ps, dd, &PullError{ErrorType: IncorrectTurnForPull}
	}
	ps.HandB = ps.HandB.Combine(ps.AttackCards.Combine(ps.DefendCards))

	dd.NumberOfCardsToDrawIntoHandA = computeDraw(ps.HandA.CardCount())
	dd.NumberOfCardsToDrawIntoHandB = computeDraw(ps.HandB.CardCount())

	return ps, dd, nil
}

func (ps PlayState) Yield() (PlayState, DrawDirective, *YieldError) {
	dd := DrawDirective{}

	if ps.AttackCards.CardCount() != ps.DefendCards.CardCount() {
		return ps, dd, &YieldError{ErrorType: IncorrectTurnForYield}
	}
	ps.Discarded = ps.Discarded.Combine(ps.AttackCards.Combine(ps.DefendCards))
	ps.AttackCards = Hand(0)
	ps.DefendCards = Hand(0)
	dd.NumberOfCardsToDrawIntoHandA = computeDraw(ps.HandA.CardCount())
	dd.NumberOfCardsToDrawIntoHandB = computeDraw(ps.HandB.CardCount())

	return ps, dd, nil

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
