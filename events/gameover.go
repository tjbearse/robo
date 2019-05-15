package events

import (
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/events/comm"
)

func StartGameWon(cc comm.CommClient, winner *game.Player) {
	c, g, err := comm.WithGameContext(cc)
	if err != nil {
		return // TODO
	}
	c.Broadcast(g, NotifyPlayerFinished{winner.Name})
}
