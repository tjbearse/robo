package loader

import (
	. "github.com/tjbearse/robo/game"
)

func GetDefaultDeck() *Deck {
	return NewDeck([]Card{
		{ 1, Forward, 1},
		{ 2, Forward, 2},
		{ 3, TurnLeft, 1},
		{ 4, TurnRight, 1},
		{ 5, Backward, 1},
	})
}
