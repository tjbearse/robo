package game

import (
	"github.com/tjbearse/robo/game/cards"
	"github.com/tjbearse/robo/game/coords"
)

type GamePhase interface {}

type Game struct {
	Board Board
	players map[*Player]bool
	Deck cards.Deck
	// Settings
	phase GamePhase
}

func NewGame(board Board, deck cards.Deck, initialPhase GamePhase) Game {
	return Game{
		board,
		map[*Player]bool{},
		deck,
		initialPhase,
	}
}

func (g Game) GetPlayers() (map[*Player]bool) {
	return g.players
}
func (g *Game) UpdatePlayers(players map[*Player]bool) {
	g.players = players
}
func (g *Game) GetPhase() GamePhase {
	return g.phase
}
func (g *Game) ChangePhase(ph GamePhase) {
	g.phase = ph
}
func (g *Game) GetNextSpawn() (coords.Configuration, error){
	return g.Board.getNextSpawn()
}

func (g *Game) CheckForRobot(c coords.Coord) *Player {
	for p, _ := range(g.players) {
		if p.Robot.Configuration != nil && p.Robot.Configuration.Location == c {
			return p
		}
	}
	return nil
}

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
