package durakGame

import (
	"durak/durakRules"
)

type GameState struct {
	durakRules.PlayState
	SuitOrder
}

type SuitOrder [4]Suit
