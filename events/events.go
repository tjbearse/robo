package events

import (
	"errors"

	"github.com/tjbearse/robo/events/comm"
)

var EventMap = map[string]func() comm.IncomingEvent {
	// Note these must be pointers for the decoding to work
	"NewGame": func () comm.IncomingEvent { return &NewGame{} },
	"JoinGame": func () comm.IncomingEvent { return &JoinGame{} },
	"LeaveGame": func () comm.IncomingEvent { return &LeaveGame{} },
	"ReadyToSpawn": func () comm.IncomingEvent { return &ReadyToSpawn{} },
	"SetSpawnHeading": func () comm.IncomingEvent { return &SetSpawnHeading{} },
	"CardToBoard": func () comm.IncomingEvent { return &CardToBoard{} },
	"CardToHand": func () comm.IncomingEvent { return &CardToHand{} },
	"CommitCards": func () comm.IncomingEvent { return &CommitCards{} },
}

var wrongPhaseError = errors.New("Not the right phase")

const RobotMaxLives int = 3
const HandSize int = 8
const MaxDamage int = HandSize
const Steps int = 5 //FIXME bad name, register count

type ErrorReport struct {
	Error string
}
