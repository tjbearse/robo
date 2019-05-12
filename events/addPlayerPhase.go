package events

import (
	"fmt"

	"github.com/tjbearse/robo/game"
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

	board := make([]*game.Card, Steps)
	r := game.Robot{name, 0, RobotMaxLives, board, nil}
	p := game.Player{name, r, game.Spawn{}, 0}
	players[&p] = true
	g.UpdatePlayers(players)
	c.Associate(&p)
	c.Broadcast(NotifyAddPlayer{name})
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
	StartSpawnPhase(c, g)
	return nil
}
