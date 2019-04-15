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

// Dir describes the facing of an object
type Dir int
// Directions, values are in numerical order clockwise
const (
	North Dir = 0
	East Dir = iota
	South
	West
)

type Coord struct {
	X int
	Y int
}

type Configuration struct {
	Location Coord
	Heading Dir
}
