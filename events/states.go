package events

import (
	"errors"

	"github.com/tjbearse/robo/game"
)

const RobotMaxLives = 3
const HandSize = 8
const Steps = 5

type IncomingEvent interface {
	Exec(commClient, *game.Game) error
}

type OutGoingEvent interface {}

type Event func() error

type commClient interface {
	Broadcast(OutGoingEvent)
	Message(OutGoingEvent, *game.Player)


	Associate(*game.Player)
	Deassociate()
	GetPlayer() *game.Player
}

func getPlayer(c commClient) (*game.Player, error) {
	p := c.GetPlayer()
	if p == nil {
		return p, errors.New("create a player first")
	}
	return p, nil
}
