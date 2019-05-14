package events

import (
	"fmt"

	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/cards"
)

type AddPlayerPhase struct {}

type JoinGame struct {
	Name string
}
// TODO restrict these actions to this phase?
// these actions are pretty general and will probably be reused
// across states. Perhaps let phases define their own behaviors
// for them and cover basics here.
func (j JoinGame) Exec(c commClient, g *game.Game) error {
	name := j.Name
	players := g.GetPlayers()
	for p := range(players) {
		if p.Name == name {
			return fmt.Errorf("name %s already exists", name)
		}
	}

	board := make([]*cards.Card, Steps)
	r := game.Robot{0, RobotMaxLives, board, nil}
	p := game.Player{name, r, game.Spawn{}, 0}
	players[&p] = true
	g.UpdatePlayers(players)
	c.Associate(&p)
	c.Message(NotifyWelcome{name}, &p)
	c.Broadcast(NotifyAddPlayer{name})
	// TODO fill in the player on what the current game state is (existing players at least)
	return nil
}

type LeaveGame struct {}
func (e LeaveGame) Exec(c commClient, g *game.Game) error {
	p, err := getPlayer(c)
	if err != nil {
		return err
	}
	players := g.GetPlayers()
	delete(players, p)
	g.UpdatePlayers(players)
	c.Deassociate()
	c.Broadcast(NotifyRemovePlayer{p.Name})
	return nil
}

type ReadyToSpawn struct {}
func (ReadyToSpawn) Exec(c commClient, g *game.Game) error {
	/*
	// TODO there should be a phase check here but the phase isn't initialized properly
	uPhase := g.GetPhase()
	_, ok := uPhase.(*AddPlayerPhase)
	_, ok2 := uPhase.(*SimulationPhase)
	if !ok && !ok2 {
		return errors.New("Not the right phase")
	}
	*/

	StartSpawnPhase(c, g)
	return nil
}
