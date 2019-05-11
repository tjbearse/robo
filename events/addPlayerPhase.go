package events

import (
	"fmt"

	"github.com/tjbearse/robo/game"
)

type AddPlayerPhase struct {}

type AddPlayer struct {
	Name string
}
func (AddPlayer) GetType() string {
	return "AddPlayer"
}

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
	p := game.Player{name, &r, game.Spawn{}}
	players[&p] = true
	g.UpdatePlayers(players)
	c.Associate(&p)
	c.Broadcast(AddPlayer{name})
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
	c.Broadcast(RemovePlayer{p.Name})
	return nil
}
type RemovePlayer struct {
	Name string
}

type ReadyToSpawn struct {}
func (ReadyToSpawn) Exec(c commClient, g *game.Game) error {
	StartSpawnPhase(c, g)
	return nil
}
