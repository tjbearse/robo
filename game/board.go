package game

import (
	"errors"
	"fmt"
)

type TileType int
const (
	Floor TileType = 0
	Pit TileType = iota
	Repair
	Upgrade
	Flag

	Conveyor
	ExpressConveyor
	Pusher
	Laser
)

type Tile struct {
	Type TileType
	Dir Dir
}

// Board represents the playing space
type Board struct {
	tiles [][]Tile
	nwalls [][]bool // n+1 x m+1
	wwalls [][]bool // n+1 x m+1
	lasers []Coord
	flags []Coord // order matters

	spawns []Configuration
	usedSpawn int
}

func NewBoard(tiles [][]Tile, nwalls [][]bool, wwalls [][]bool, flags []Coord, spawns []Configuration) (*Board, error) {
	if len(tiles) == 0 ||
		len(tiles[0]) == 0 ||
		len(tiles[0]) + 1 != len(nwalls[0]) ||
		len(tiles) + 1 != len(nwalls) ||
		len(tiles[0]) + 1 != len(wwalls[0]) ||
		len(tiles) + 1 != len(wwalls) {
		return nil, errors.New("Invalid tile or wall array size")
	}
	// TODO maybe relax this a bit
	if len(spawns) < 8 {
		return nil, errors.New("Less than 8 spawns")
	}


	lasers := make([]Coord, 0)
	for i, row := range(tiles) {
		for j, tile := range(row) {
			if tile.Type == Laser {
				lasers = append(lasers, Coord{i,j})
			}
		}
	}

	// TODO validate flags are indeed flags and the only flags

	board := Board{tiles, nwalls, wwalls, lasers, flags, spawns, 0}

	for _, spawn := range(spawns) {
		if !board.IsInbounds(spawn.Location) {
			return nil, errors.New("Spawn out of range")
		}
		tile, _ := board.GetTile(spawn.Location)
		if tile.Type != Floor {
			return nil, errors.New("Spawn must be a floor")
		}
	}
	return &board, nil
}

func (b *Board) GetNumFlags() (int) {
	return len(b.flags)
}

func (b *Board) GetFlag(i int) (Coord, error) {
	if i < 0 || i >= len(b.flags) {
		return Coord{}, fmt.Errorf("out of range index: %d", i)
	}
	return b.flags[i], nil
}

func (b *Board) GetTile(c Coord) (Tile, error) {
	if !b.IsInbounds(c) {
		return Tile{}, errors.New("Out of bounds")
	}
	return b.tiles[c.X][c.Y], nil
}

func (b *Board) IsInbounds(c Coord) bool {
	width := len(b.tiles)
	height := len(b.tiles[0]) // FIXME
	return c.X >= 0 && c.X < width &&
		c.Y >= 0 && c.Y < height
}

func (b *Board) IsPassable(from Coord, to Coord) (bool, error) {
	dx := abs(to.X - from.X)
	dy := abs(to.Y - from.Y)
	if (dx != 0 && dy != 0) || (dx != 1 && dy != 1) {
		return false, fmt.Errorf("not a length 1 offset from %v to %v",  to, from)
	}
	return !b.wallBetween(to, from), nil
}

// needs offset length 1
func (board *Board) wallBetween(a Coord, b Coord) bool {
	offset := a.OffsetTo(b)
	if offset.X != 0 {
		// check in W walls
		x := max(a.X, b.X)
		if x > len(board.wwalls) || a.Y > len(board.wwalls[x]) {
			return false
		}
		return board.wwalls[x][a.Y]
	} else {
		// check in N walls
		y := max(a.Y, b.Y)
		if a.X > len(board.nwalls) || y > len(board.nwalls[a.X]) {
			return false
		}
		return board.nwalls[a.X][y]
	}
}

func abs (a int) int {
	return max(a, -a)
}
func max (a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (b *Board) getNextSpawn() (Configuration, error) {
	if b.usedSpawn == len(b.spawns) {
		return Configuration{}, errors.New("no more spawns")
	}
	c := b.spawns[b.usedSpawn]
	b.usedSpawn += 1
	return c, nil
}

