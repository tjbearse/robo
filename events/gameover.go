package events

import (
	"github.com/tjbearse/robo/game"
)

func StartGameWon(c commClient, g *game.Game, winner *game.Player) {
	c.Broadcast(NotifyPlayerFinished{winner.Name})
}
