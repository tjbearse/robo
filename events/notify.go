package events

import (
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/cards"
	"github.com/tjbearse/robo/game/coords"
)

// Events that notify the players of happenings
type NotifyBoard struct {
	Board game.PlainBoard
}

type NotifyNewGame struct {
	GameId GameId
}

type NotifyRemovePlayer struct {
	Name string
}

type NotifyAddPlayer struct {
	Name string
}
type NotifyWelcome struct {
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
	OldConfig coords.Configuration
	NewConfig coords.Configuration
}

// FIXME clarity between robot and player names
// is there a better way to identify, e.g. a typed ID?
type NotifyRobotFell struct {
	Name string
	Reason MoveReason
	OldConfig coords.Configuration
	Target coords.Configuration
}

type NotifyRevealCard struct {
	Name string
	Card cards.Card
}

type NotifySpawnUpdate struct {
	Name string
	Coord coords.Coord
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
	Card cards.Card
}

type NotifyCardToHand struct {
	BoardSlot uint
	HandOffset uint
	Card cards.Card
}

type NotifyCardToBoardBlind struct {
	Player string
	BoardSlot uint
}

type NotifyCardToHandBlind struct {
	Player string
	BoardSlot uint
}

// FIXME too generic
type NotifyPlayerReady struct {
	Name string
}
