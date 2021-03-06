package loader

import (
	. "github.com/tjbearse/robo/game/cards"
)

// TODO move this to a config option
func GetDefaultDeck() (*Deck, error) {
	return NewDeck([]Card{
		{ 10, UTurn, 0 },
		{ 20, UTurn, 0 },
		{ 30, UTurn, 0 },
		{ 40, UTurn, 0 },
		{ 50, UTurn, 0 },
		{ 60, UTurn, 0 },
		{ 70, RotateLeft, 0 },
		{ 80, RotateRight, 0 },
		{ 90, RotateLeft, 0 },
		{ 100, RotateRight, 0 },
		{ 110, RotateLeft, 0 },
		{ 120, RotateRight, 0 },
		{ 130, RotateLeft, 0 },
		{ 140, RotateRight, 0 },
		{ 150, RotateLeft, 0 },
		{ 160, RotateRight, 0 },
		{ 170, RotateLeft, 0 },
		{ 180, RotateRight, 0 },
		{ 190, RotateLeft, 0 },
		{ 200, RotateRight, 0 },
		{ 210, RotateLeft, 0 },
		{ 220, RotateRight, 0 },
		{ 230, RotateLeft, 0 },
		{ 240, RotateRight, 0 },
		{ 250, RotateLeft, 0 },
		{ 260, RotateRight, 0 },
		{ 270, RotateLeft, 0 },
		{ 280, RotateRight, 0 },
		{ 290, RotateLeft, 0 },
		{ 300, RotateRight, 0 },
		{ 310, RotateLeft, 0 },
		{ 320, RotateRight, 0 },
		{ 330, RotateLeft, 0 },
		{ 340, RotateRight, 0 },
		{ 350, RotateLeft, 0 },
		{ 360, RotateRight, 0 },
		{ 370, RotateLeft, 0 },
		{ 380, RotateRight, 0 },
		{ 390, RotateLeft, 0 },
		{ 400, RotateRight, 0 },
		{ 410, RotateLeft, 0 },
		{ 420, RotateRight, 0 },
		{ 430, BackUp, 1 },
		{ 440, BackUp, 1 },
		{ 450, BackUp, 1 },
		{ 460, BackUp, 1 },
		{ 470, BackUp, 1 },
		{ 480, BackUp, 1 },
		{ 490, Move, 1, },
		{ 500, Move, 1, },
		{ 510, Move, 1, },
		{ 520, Move, 1, },
		{ 530, Move, 1, },
		{ 540, Move, 1, },
		{ 550, Move, 1, },
		{ 560, Move, 1, },
		{ 570, Move, 1, },
		{ 580, Move, 1, },
		{ 590, Move, 1, },
		{ 600, Move, 1, },
		{ 610, Move, 1, },
		{ 620, Move, 1, },
		{ 630, Move, 1, },
		{ 640, Move, 1, },
		{ 650, Move, 1, },
		{ 660, Move, 1, },
		{ 670, Move, 2, },
		{ 680, Move, 2, },
		{ 690, Move, 2, },
		{ 700, Move, 2, },
		{ 710, Move, 2, },
		{ 720, Move, 2, },
		{ 730, Move, 2, },
		{ 740, Move, 2, },
		{ 750, Move, 2, },
		{ 760, Move, 2, },
		{ 770, Move, 2, },
		{ 780, Move, 2, },
		{ 790, Move, 3, },
		{ 800, Move, 3, },
		{ 810, Move, 3, },
		{ 820, Move, 3, },
		{ 830, Move, 3, },
		{ 840, Move, 3, },
	})
}
