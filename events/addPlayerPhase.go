package events

import (
	"github.com/tjbearse/robo/game"
)

type AddPlayerPhase struct {}

type Event func () error

func AddPlayer(g *game.Game, name string) Event {
	return func () error {
		board := make([]*game.Card, 0, Steps)
		r := game.Robot{name, 0, RobotMaxLives, board, nil}
		p := game.Player{name, r, nil}
		players := g.GetPlayers()
		players[&p] = true
		g.UpdatePlayers(players)
		return nil
	}
}

// FIXME should events look up players, e.g. by name?
func RemovePlayer(g *game.Game, p *game.Player) Event {
	return func () error {
		players := g.GetPlayers()
		delete(players, p)
		g.UpdatePlayers(players)
		return nil
	}
}

func Ready(g *game.Game) Event {
	return func () error {
		ph := NewSpawnPhase(g)
		g.ChangePhase(ph)
		return nil
	}
}
