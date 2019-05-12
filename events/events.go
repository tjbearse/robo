package events

import (
	"github.com/tjbearse/robo/game"
)

type IncomingEvent interface {
	Exec(commClient, *game.Game) error
}

type OutGoingEvent interface {}
