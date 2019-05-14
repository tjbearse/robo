import {Card} from './card'
import {Coord, Config} from './coord'

export interface Player {
	name: string
	robot: Robot
	hand: Card[]
	board: { [i: number]: Card } // TODO flipped card
	flagNum: number
	spawn: Coord
}

export function newPlayer(name:string) : Player {
	return {
		name: name,
		hand: [],
		board: {},
		robot: newRobot(),
		flagNum: 0,
		spawn: null,
	}
}

function newRobot() : Robot {
	return {
		damage: 0,
		lives: 0,
		config: null,
	}
}

export interface Robot {
	damage: number
	lives: number
	config: Config
}
