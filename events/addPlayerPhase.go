package events

import (
	"errors"
	"fmt"

	"github.com/tjbearse/robo/events/comm"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/cards"
)

type AddPlayerPhase struct {}

// TODO need to be able to handle zero state everywhere!
type JoinGame struct {
	PlayerName string
	Game GameId
}
// TODO restrict these actions to this phase?
// these actions are pretty general and will probably be reused
// across states. Perhaps let phases define their own behaviors
// for them and cover basics here.
func (j JoinGame) Exec(cc comm.CommClient) error {
	c,err := comm.WithoutContext(cc)
	if err != nil {
		return err
	}

	g, ok := gameStore[j.Game]
	if !ok {
		return fmt.Errorf("game (%s) doesn't exist", j.Game)
	}

	if j.PlayerName == "" {
		return errors.New("name cannot be empty")
	}
	name := j.PlayerName
	players := g.GetPlayers()
	for p := range(players) {
		if p.Name == name {
			return fmt.Errorf("name %s already exists", name)
		}
	}

	board := make([]*cards.Card, Steps)
	r := game.Robot{0, RobotMaxLives, board, nil}
	p := &game.Player{name, r, game.SpawnSetting{}, 0}
	players[p] = true
	g.UpdatePlayers(players)
	c.SetGame(g)
	c.SetPlayer(p)
	c.Reply(NotifyWelcome{j.Game, name})
	// fill in player on other existing players
	for op := range(players) {
		if op != p {
			c.Reply(NotifyAddPlayer{op.Name})
		}
	}
	c.Broadcast(g, NotifyAddPlayer{name})
	return nil
}

type LeaveGame struct {}
func (e LeaveGame) Exec(cc comm.CommClient) error {
	c, g, p, err := comm.WithPlayerContext(cc)
	if err != nil {
		return err
	}
	players := g.GetPlayers()
	delete(players, p)
	g.UpdatePlayers(players)
	c.Clear()
	c.Broadcast(g, NotifyRemovePlayer{p.Name})
	return nil
}

type ReadyToSpawn struct {}
func (ReadyToSpawn) Exec(cc comm.CommClient) error {
	c, g, err := comm.WithGameContext(cc)
	if err != nil {
		return err
	}
	uPhase := g.GetPhase()
	if _, ok := uPhase.(*AddPlayerPhase); ok {
		c.Broadcast(g, NotifyBoard{g.Board.GetBoardDump()})
	} else if _, ok = uPhase.(*SimulationPhase); !ok {
		return wrongPhaseError
	}

	StartSpawnPhase(cc)
	return nil
}
