package game

import (
	"errors"
	"fmt"

	"github.com/tjbearse/robo/game/coords"
)

type TileType int
const (
	Floor TileType = 0
	Pit TileType = iota
	Repair
	Upgrade
	Flag
	Spawn

	Conveyor
	ExpressConveyor
	Pusher
	Gear // TODO this needs some more context re: direction
	Laser
)

type Tile struct {
	Type TileType
	Dir coords.Dir
}

type PlainBoard struct {
	Tiles [][]Tile
	Nwalls [][]bool
	Wwalls [][]bool
	FlagOrder []coords.Coord
}
// Board represents the playing space
type Board struct {
	tiles [][]Tile
	// FIXME nwall only needs to be one taller, not one wider than Tiles
	nwalls [][]bool // x+1, y+1 -> x,y+1
	// FIXME similar for wwalls but reversed
	wwalls [][]bool // x+1, y+1 -> x+1,y
	lasers []coords.Coord
	flagOrder []coords.Coord

	spawns []coords.Configuration
	usedSpawn int
}

func NewBoard(plain PlainBoard, spawns []coords.Configuration) (*Board, error) {
	tiles := plain.Tiles
	nwalls := plain.Nwalls
	wwalls := plain.Wwalls
	flags := plain.FlagOrder
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


	lasers := make([]coords.Coord, 0)
	for i, row := range(tiles) {
		for j, tile := range(row) {
			if tile.Type == Laser {
				lasers = append(lasers, coords.Coord{i,j})
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
		if tile.Type != Spawn {
			return nil, errors.New("Spawn must be a spawn Tile")
		}
	}
	return &board, nil
}

func (b *Board) GetBoardDump() PlainBoard {
	return PlainBoard{
		b.tiles,
		b.nwalls,
		b.wwalls,
		b.flagOrder,
	}
}

func (b *Board) GetNumFlags() (int) {
	return len(b.flagOrder)
}

func (b *Board) GetFlag(i int) (coords.Coord, error) {
	if i < 0 || i >= len(b.flagOrder) {
		return coords.Coord{}, fmt.Errorf("out of range index: %d", i)
	}
	return b.flagOrder[i], nil
}

func (b *Board) GetTile(c coords.Coord) (Tile, error) {
	if !b.IsInbounds(c) {
		return Tile{}, errors.New("Out of bounds")
	}
	return b.tiles[c.X][c.Y], nil
}

func (b *Board) IsInbounds(c coords.Coord) bool {
	width := len(b.tiles)
	height := len(b.tiles[0]) // FIXME
	return c.X >= 0 && c.X < width &&
		c.Y >= 0 && c.Y < height
}

func (b *Board) IsPassable(from coords.Coord, to coords.Coord) (bool, error) {
	dx := abs(to.X - from.X)
	dy := abs(to.Y - from.Y)
	if (dx != 0 && dy != 0) || (dx != 1 && dy != 1) {
		return false, fmt.Errorf("not a length 1 offset from %v to %v",  to, from)
	}
	return !b.wallBetween(to, from), nil
}

// needs offset length 1
func (board *Board) wallBetween(a coords.Coord, b coords.Coord) bool {
	offset := a.OffsetTo(b)
	if offset.X != 0 {
		// check in W walls
		x := max(a.X, b.X)
		if x >= len(board.wwalls) || a.Y >= len(board.wwalls[x]) {
			return false
		}
		return board.wwalls[x][a.Y]
	} else {
		// check in N walls
		y := max(a.Y, b.Y)
		if a.X >= len(board.nwalls) || y >= len(board.nwalls[a.X]) {
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

func (b *Board) getNextSpawn() (coords.Configuration, error) {
	if b.usedSpawn == len(b.spawns) {
		return coords.Configuration{}, errors.New("no more spawns")
	}
	c := b.spawns[b.usedSpawn]
	b.usedSpawn += 1
	return c, nil
}

func (b *Board) GetLasers() []coords.Coord {
	return b.lasers
}
