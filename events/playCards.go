package events

import (
	"errors"
	"fmt"

	"github.com/tjbearse/robo/game"
)


// TODO there is an option to power down instead of receiving cards
func StartCardsPhase (c commClient, g *game.Game) {
	hands := map[*game.Player][]game.Card{}
	players := g.GetPlayers()
	for p := range(players) {
		if p.Robot.Lives > 0 {
			cards := g.Deck.Deal(HandSize - p.Robot.Damage)
			hands[p] = cards
			prompt := PromptWithHand{hands[p]}
			c.Message(prompt, p)
		}
	}
	newPh := PlayCardsPhase{hands, map[*game.Player]bool{}}
	g.ChangePhase(&newPh)
}

type PlayCardsPhase struct {
	Hands map[*game.Player][]game.Card
	Ready map[*game.Player]bool
	// TODO: Timer chan bool
}


type CardToBoard struct {
	HandOffset uint
	BoardSlot uint
}

func (e CardToBoard) Exec(c commClient, g *game.Game) error {
	p, err := getPlayer(c)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*PlayCardsPhase)
	if !ok {
		return errors.New("Not the right phase")
	}

	hand := ph.Hands[p]
	co := e.HandOffset
	if int(co) > len(hand) {
		return fmt.Errorf("card out of range: %d", co)
	}
	card := hand[co]

	slot := e.BoardSlot
	if int(slot) >= len(p.Robot.Board) {
		return fmt.Errorf("slot out of range: %d", slot)
	}
	if p.Robot.Board[slot] != nil {
		return fmt.Errorf("slot not free: %d", slot)
	}

	p.Robot.Board[slot] = &card
	hand = append(hand[:co], hand[co+1:]...)
	ph.Hands[p] = hand

	c.Broadcast(NotifyCardToBoardBlind{p.Name, slot})
	c.Message(NotifyCardToBoard{slot, co, card}, p)
	return nil
}

type CardToHand struct {
	BoardSlot uint
}
func (e CardToHand) Exec(c commClient, g *game.Game) error {
	p, err := getPlayer(c)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*PlayCardsPhase)
	if !ok {
		// TODO abstract this repetitive error message
		return errors.New("Not the right phase")
	}

	slot := e.BoardSlot
	if int(slot) > len(p.Robot.Board) {
		return fmt.Errorf("slot out of range: %d", slot)
	} else if p.Robot.Board[slot] == nil {
		return fmt.Errorf("slot not occupied: %d", slot)
	} else if int(slot) > (HandSize - p.Robot.Damage) {
		return fmt.Errorf("slot locked: %d", slot)
	}
	hand := ph.Hands[p]
	co := uint(len(hand))
	card := *p.Robot.Board[slot]
	ph.Hands[p] = append(hand, card)
	p.Robot.Board[slot] = nil

	c.Broadcast(NotifyCardToHandBlind{p.Name, slot})
	c.Message(NotifyCardToHand{slot, co, card}, p)
	return nil
}

type CommitCards struct {}
func (e CommitCards) Exec(c commClient, g *game.Game) error {
	p, err := getPlayer(c)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*PlayCardsPhase)
	if !ok {
		return errors.New("Not the right phase")
	}
	// TODO check if hand full
	if ph.Ready[p] == true {
		return errors.New("Already ready")
	}
	ph.Ready[p] = true
	c.Broadcast(NotifyPlayerReady{p.Name})

	for _, ready := range(ph.Ready) {
		if !ready {
			return nil
		}
	}
	// all ready
	StartSimulationPhase(c,g)
	return nil
}
