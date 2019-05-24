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
	GameId GameId
	Name string
}


// Reasons for movement
type MoveReason string
const (
	Spawned MoveReason = "spawned"
	Moved = "moved"
	Bumped = "bumped"
	Conveyed = "conveyed"
	Gear = "gear"
)

type NotifyRobotMoved struct {
	Player string
	Reason MoveReason
	OldConfig coords.Configuration
	NewConfig coords.Configuration
}

// TODO is there a better way to identify players, e.g. a typed ID?
type NotifyRobotFell struct {
	Player string
	Reason MoveReason
	OldConfig coords.Configuration
	Target coords.Configuration
}

type NotifyRevealCard struct {
	Player string
	Card cards.Card
}

type NotifySpawnUpdate struct {
	Player string
	Coord coords.Coord
}

type NotifyFlagTouched struct {
	Player string
	FlagNum int
}

type NotifyPlayerFinished struct {
	Player string
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

type NotifyRandomBoardFill struct {
	Player string
	BoardSlots []uint
}

type NotifyCardToHandBlind struct {
	Player string
	BoardSlot uint
}

type NotifyCleanup struct {
	Player string
	Board []*cards.Card
}

// FIXME too generic
type NotifyPlayerReady struct {
	Player string
}

type Ray struct {
	Start coords.Coord
	StartPlayer string
	End coords.Coord
	EndPlayer string
}
type NotifyLaserFire struct {
	Rays []Ray
}

type NotifyDamage struct {
	Player string
	NewTotal int
}

type NotifyHeal struct {
	Player string
	NewTotal int
}

type NotifyLifeLoss struct {
	Player string
	NewTotal int
}
