package game

import (
	"errors"
)

// Board represents the playing space
type Board struct {
	// Tiles with true are floor, false are pits
	Tiles [][]bool
	// Hazards []Hazard
	// Entities []Entities
	spawns []Configuration
	usedSpawn int
}

func (b *Board) getNextSpawn() (Configuration, error) {
	if b.usedSpawn == len(b.spawns) {
		return Configuration{}, errors.New("no more spawns")
	}
	c := b.spawns[b.usedSpawn]
	b.usedSpawn += 1
	return c, nil
}


type Coord struct {
	X int
	Y int
}

type Configuration struct {
	Location Coord
	Heading Dir
}
