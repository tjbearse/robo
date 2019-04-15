package events

import (
	"errors"

	"github.com/tjbearse/robo/game"
)

func NewPlayCardsPhase (g *game.Game) PlayCardsPhase {
	hands := map[*game.Player][]game.Card{}
	players := g.GetPlayers()
	for p := range(players) {
		if p.Robot.Lives > 0 {
			hands[p] = g.Deck.Deal(HandSize - p.Robot.Damage)
			// TODO prompt with hand
			// FIXME move hands into player objects?, i.e. update player
		}
	}
	return PlayCardsPhase{hands, map[*game.Player]bool{}}
}


type PlayCardsPhase struct {
	Hands map[*game.Player][]game.Card
	Ready map[*game.Player]bool
	// TODO: Timer chan bool
}
	// State:
	// cards in player hand
	// cards committed
	// timer state

	// Actions:
	// start timer
	// add selection
	// unset selection
	// ready/submit
	// (out) cancelled by timer, random vals

	// GOTO
	// simulatePhase when all spawned

func AddSelection(g *game.Game, p *game.Player, c int, slot int) Event {
	return func () error {
		uPhase := g.GetPhase()
		ph, ok := uPhase.(PlayCardsPhase)
		if !ok {
			return errors.New("Not the right phase")
		}
		hand := ph.Hands[p]
		// check above ^
		if c > len(hand) {
			return errors.New("card out of range")
		}
		card := hand[c]
		if slot > len(p.Robot.Board) {
			return errors.New("slot out of range")
		}
		if p.Robot.Board[slot] != nil {
			return errors.New("slot not free")
		}
		// FIXME use the update pattern
		p.Robot.Board[slot] = &card
		ph.Hands[p] = append(hand[:c], hand[c+1:]...)
		// TODO updated robot board and hand
		return nil
	}
}

// TODO convert to event style
func (ph *PlayCardsPhase) removeSelection(g *game.Game, p *game.Player, slot int) {
	hand := ph.Hands[p]
	// TODO check above ^
	if slot > len(p.Robot.Board) {
		// TODO swat: slot out of range
	}
	if p.Robot.Board[slot] == nil {
		// TODO swat: slot not free
	}
	ph.Hands[p] = append(hand, *p.Robot.Board[slot])
	p.Robot.Board[slot] = nil
	// TODO update player board and hand updated
}
