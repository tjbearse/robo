package events

import (
	"github.com/tjbearse/robo/game"
)

// Events that notify the players of happenings

type NotifyRemovePlayer struct {
	Name string
}

type NotifyAddPlayer struct {
	Name string
}


// Reasons for movement
type MoveReason string
const (
	Spawned MoveReason = "spawned"
	Moved = "moved"
	Bumped = "bumped"
)

type NotifyRobotMoved struct {
	Name string
	Reason MoveReason
	OldConfig game.Configuration
	NewConfig game.Configuration
}

// FIXME clarity between robot and player names
// is there a better way to identify, e.g. a typed ID?
type NotifyRobotFell struct {
	Name string
	Reason MoveReason
	OldConfig game.Configuration
	Target game.Configuration
}

type NotifyRevealCard struct {
	Name string
	Card game.Card
}

type NotifySpawnUpdate struct {
	Name string
	Coord game.Coord
}

type NotifyFlagTouched struct {
	Name string
	FlagNum int
}

type NotifyPlayerFinished struct {
	Name string
}

type NotifyStartSpawn struct {}

type NotifyCardToBoard struct {
	BoardSlot uint
	HandOffset uint
	Card game.Card
}

type NotifyCardToHand struct {
	BoardSlot uint
	HandOffset uint
	Card game.Card
}

type NotifyCardToBoardBlind struct {
	Player string
	BoardSlot uint
}
type NotifyCardToHandBlind struct {
	Player string
	BoardSlot uint
}

type NotifyPlayerReady struct {
	Name string
}
