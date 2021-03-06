package events

import (
	"github.com/tjbearse/robo/game/cards"
	"github.com/tjbearse/robo/game/coords"
)

// Events that request player input
// TODO consolidate into notify?

type PromptForSpawn struct {
	Player string
	Coord coords.Coord
}

type PromptWithHand struct {
	Cards []cards.Card
}
