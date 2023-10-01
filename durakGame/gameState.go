package durakGame

import (
	"github.com/Ivaan/durak/durakRules"
)

type GameState struct {
	PlayerA  durakRules.Hand
	PlayerB  durakRules.Hand
	DrawDeck []int //bit possition of card, yes this is a tighter coupling than I imagined .. this is fine
	SuitOrder
	PlayerBAttacker   bool //otherwise PlayerA is PlayerBAttacker
	DefenderDefending bool //otherwise PlayerBAttacker attacking
}

type SuitOrder [4]Suit
