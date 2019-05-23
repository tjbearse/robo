package loader

import (
	. "github.com/tjbearse/robo/game"
	. "github.com/tjbearse/robo/game/coords"
)

func GetDefaultBoard() (*Board, error) {
	width, height := 10, 10
	pb := PlainBoard{}
	pb.Tiles = make([][]Tile, width)
	for x := 0; x < width; x++ {
		pb.Tiles[x] = make([]Tile, height)
	}
	types := []TileType{Pit, Repair, Upgrade, Conveyor, ExpressConveyor, Pusher, Laser}
	for i := 0; i < width && i < height; i++ {
		idx := i % len(types)
		pb.Tiles[i][i].Type = types[idx]
	}

	pb.Wwalls = make([][]bool, width+1)
	for x := 0; x < len(pb.Wwalls); x++ {
		pb.Wwalls[x] = make([]bool, height+1)
	}
	pb.Nwalls = make([][]bool, width+1)
	for x := 0; x < len(pb.Nwalls); x++ {
		pb.Nwalls[x] = make([]bool, height+1)
	}
	pb.Nwalls[3][4] = true
	pb.Wwalls[8][2] = true

	pb.Tiles[2][4].Type = Flag
	pb.FlagOrder = []Coord{{2,4}}

	spawns := []Configuration{
		{Coord{3,4},East},
		{Coord{0,1},East},
		{Coord{0,2},East},
		{Coord{0,3},East},
		{Coord{0,4},East},
		{Coord{0,5},East},
		{Coord{0,6},East},
		{Coord{0,7},East},
	}
	return NewBoard(pb, spawns)
}
