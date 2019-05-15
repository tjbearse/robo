package events

import (
	"errors"
	"fmt"

	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/cards"
	"github.com/tjbearse/robo/events/comm"
)


// TODO there is an option to power down instead of receiving cards
func StartCardsPhase (cc comm.CommClient) {
	c, g, err := comm.WithGameContext(cc)
	if err != nil {
		return // TODO
	}
	hands := map[*game.Player][]cards.Card{}
	players := g.GetPlayers()
	for p := range(players) {
		if p.Robot.Lives > 0 {
			cards := g.Deck.Deal(HandSize - p.Robot.Damage)
			hands[p] = cards
			// TODO add info for num cards of others?
			prompt := PromptWithHand{hands[p]}
			c.MessagePlayer(p, prompt)
		}
	}
	newPh := PlayCardsPhase{hands, map[*game.Player]bool{}}
	g.ChangePhase(&newPh)
}

type PlayCardsPhase struct {
	Hands map[*game.Player][]cards.Card
	Ready map[*game.Player]bool
	// TODO: Timer chan bool
}


type CardToBoard struct {
	HandOffset uint
	BoardSlot uint
}

func (e CardToBoard) Exec(cc comm.CommClient) error {
	c, g, p, err := comm.WithPlayerContext(cc)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*PlayCardsPhase)
	if !ok {
		return wrongPhaseError
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

	c.Broadcast(g, NotifyCardToBoardBlind{p.Name, slot})
	c.MessagePlayer(p, NotifyCardToBoard{slot, co, card})
	return nil
}

type CardToHand struct {
	BoardSlot uint
}
func (e CardToHand) Exec(cc comm.CommClient) error {
	c, g, p, err := comm.WithPlayerContext(cc)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*PlayCardsPhase)
	if !ok {
		return wrongPhaseError
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

	c.Broadcast(g, NotifyCardToHandBlind{p.Name, slot})
	c.MessagePlayer(p, NotifyCardToHand{slot, co, card})
	return nil
}

type CommitCards struct {}
func (e CommitCards) Exec(cc comm.CommClient) error {
	c, g, p, err := comm.WithPlayerContext(cc)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*PlayCardsPhase)
	if !ok {
		return wrongPhaseError
	}
	// TODO check if actually played cards
	if ph.Ready[p] == true {
		return errors.New("Already ready")
	}
	ph.Ready[p] = true
	c.Broadcast(g, NotifyPlayerReady{p.Name})
	for _, card := range(ph.Hands[p]) {
		g.Deck.Discard(card)
	}

	for _, ready := range(ph.Ready) {
		if !ready {
			return nil
		}
	}
	// all ready
	StartSimulationPhase(cc)
	return nil
}
