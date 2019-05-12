package events

import (
	"github.com/tjbearse/robo/game"
)

// Events that request player input

type PromptForSpawn struct {
	Robot string
	Location game.Coord
}

type PromptWithHand struct {
	Cards []game.Card
}
