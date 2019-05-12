package events

import (
	"errors"

	"github.com/tjbearse/robo/game"
)

// TODO rm this file once game logic removed


const RobotMaxLives int = 3
const HandSize int = 8
const Steps int = 5

type commClient interface {
	Broadcast(OutGoingEvent)
	Message(OutGoingEvent, *game.Player)


	Associate(*game.Player)
	Deassociate()
	GetPlayer() *game.Player
}

// TODO get rid of this fn? Just share an error message if player is nil
func getPlayer(c commClient) (*game.Player, error) {
	p := c.GetPlayer()
	if p == nil {
		return p, errors.New("create a player first")
	}
	return p, nil
}
