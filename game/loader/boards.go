package loader

import (
	. "github.com/tjbearse/robo/game"
)


func GetDefaultBoard() (*Board, error) {
	width, height := 10, 10
	tiles := make([][]Tile, width)
	for x := 0; x < width; x++ {
		tiles[x] = make([]Tile, height)
	}
	types := []TileType{Pit, Repair, Upgrade, Conveyor, ExpressConveyor, Pusher, Laser}
	for i := 0; i < width && i < height; i++ {
		idx := i % len(types)
		tiles[i][i].Type = types[idx]
	}

	wwalls := make([][]bool, width+1)
	for x := 0; x < len(wwalls); x++ {
		wwalls[x] = make([]bool, height+1)
	}
	nwalls := make([][]bool, width+1)
	for x := 0; x < len(nwalls); x++ {
		nwalls[x] = make([]bool, height+1)
	}

	tiles[2][4].Type = Flag
	flags := []Coord{{2,4}}

	spawns := []Configuration{
		{Coord{1,0},West},
		{Coord{0,1},West},
		{Coord{0,2},West},
		{Coord{0,3},West},
		{Coord{0,4},West},
		{Coord{0,5},West},
		{Coord{0,6},West},
		{Coord{0,7},West},
	}
	return NewBoard(tiles, nwalls, wwalls, flags, spawns)
}
