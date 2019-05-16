import { createReducer } from 'redux-starter-kit'
import notify from '../actions/notify'
import { Tile, Board, Walls } from '../types/board'
import { Coord } from '../types/coord'

const boardReducer = createReducer([], {
    [notify.Board]: (state, action) => {
		let board : Tile[][] = []
		let { Tiles, Nwalls, Wwalls, FlagOrder} = action.payload.Board

		for (let x=0; x < Tiles.length; x++ ) {
			board.push([])
			for (let y=0; y < Tiles[x].length; y++) {
				let t = Tiles[x][y]
				let nwall = Nwalls[x][y]
				let swall = Nwalls[x][y+1]
				let wwall = Wwalls[x][y]
				let ewall = Wwalls[x+1][y]
				let tile = getTile(t, nwall, ewall, swall, wwall)
				board[x].push(tile)
			}
		}
		for (let i=0; i < FlagOrder.length; i++) {
			let coord : Coord = FlagOrder[i]
			board[coord.X][coord.Y].num = i
		}

		return board
	},
	[notify.Welcome]: (state, action) => [],
})

export default boardReducer

//--

function getTile(serverTile, nWall, eWall, sWall, wWall) : Tile {
	// Convert walls
	let wall : Walls = Walls.None
	let conv = [
		[nWall, Walls.North],
		[eWall, Walls.East],
		[sWall, Walls.South],
		[wWall, Walls.West],
	]
	conv.forEach(([exists, key]) => {
		if (exists) {
			wall |= key
		}
	})

	return {
		type: serverTile.Type,
		dir: serverTile.Dir,
		walls: wall,
	}
}
/*

/*
  Tiles [][]Tile
  Nwalls [][]bool
  Wwalls [][]bool
  FlagOrder []coords.Coord
*/
