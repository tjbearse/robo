package game

// Board represents the playing space
type Board struct {
	// Tiles with true are floor, false are pits
	Tiles [][]bool
	Robots []Robot
	// Hazards []Hazard
	// Entities []Entities
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
