package game

import (
	"github.com/tjbearse/robo/game/cards"
	"github.com/tjbearse/robo/game/coords"
)

// TODO collapse player and Robot?
type Player struct {
	Name string
	Robot Robot
	Spawn SpawnSetting
	FlagNum int // i.e. the flag currently targetting
}

type Robot struct {
	Damage int
	Lives int
	Board []*cards.Card
	Configuration *coords.Configuration
}

type SpawnState int
const (
	Unset SpawnState = iota
	Rotatable
)
type SpawnSetting struct {
	State SpawnState
	Coord coords.Coord
}
