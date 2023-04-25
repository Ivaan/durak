package durakRules

import "fmt"

type AttackErrorType int

const (
	IncorrectTurnForAttack AttackErrorType = iota
	NotExactlyOneAttackCard
	HandANotHoldingCard
	CardNotOnTable
)

type AttackError struct {
	ErrorType AttackErrorType
	*PlayState
	Play *Hand
}

func (e *AttackError) Error() string {
	switch e.ErrorType {
	case IncorrectTurnForAttack:
		return fmt.Sprintf("Incorrect turn for attack")
	case NotExactlyOneAttackCard:
		return fmt.Sprintf("Must attack with exacty one card, attacked with %s.", e.Play.GetCards())
	case HandANotHoldingCard:
		return fmt.Sprintf("Hand A is not holding the card %s.", e.Play.GetCards())
	case CardNotOnTable:
		return fmt.Sprintf("Rank of card %s is not already on the table: %s.", e.Play.GetCards(), e.PlayState)
	default:
		return fmt.Sprintf("Unrecognized error %d", int(e.ErrorType))
	}
}

type DefendErrorType int

const (
	IncorrectTurnForDefend DefendErrorType = iota
	NotExactlyOneDefedCard
	HandBNotHoldingCard
	NotAllCardsBeaten
)

type DefendError struct {
	ErrorType DefendErrorType
	*PlayState
	Play *Hand
}

func (e *DefendError) Error() string {
	switch e.ErrorType {

	case IncorrectTurnForDefend:
		return fmt.Sprintf("Incorrect turn for defend")
	case HandBNotHoldingCard:
		return fmt.Sprintf("Hand B is not holding the card %s", e.Play.GetCards())
	case NotAllCardsBeaten:
		return fmt.Sprintf("Not all attack cards beaten")
	default:
		return fmt.Sprintf("Unrecognized error %d", int(e.ErrorType))
	}
}

type PullErrorType int

const (
	IncorrectTurnForPull PullErrorType = iota
)

type PullError struct {
	ErrorType PullErrorType
	*PlayState
	Play *Hand
}

func (e *PullError) Error() string {
	switch e.ErrorType {
	case IncorrectTurnForPull:
		return fmt.Sprintf("Incorrect turn for pull")
	default:
		return fmt.Sprintf("Unrecognized error %d", int(e.ErrorType))
	}
}

type YieldErrorType int

const (
	IncorrectTurnForYield YieldErrorType = iota
)

type YieldError struct {
	ErrorType YieldErrorType
	*PlayState
	Play *Hand
}

func (e *YieldError) Error() string {
	switch e.ErrorType {
	case IncorrectTurnForYield:
		return fmt.Sprintf("Incorrect turn for yield")
	default:
		return fmt.Sprintf("Unrecognized error %d", int(e.ErrorType))
	}
}
