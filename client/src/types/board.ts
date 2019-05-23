export enum Directions {
	North = "North",
	East = "East",
	South = "South",
	West = "West",
}

export enum Walls {
	None  = 0,
	North = 1 << 0,
	East  = 1 << 1,
	South = 1 << 2,
	West  = 1 << 3,
}

export enum TileType {
	// order matters to match with server
	Floor = 0,
	Pit,
	Repair,
	Upgrade,
	Flag,
	Spawn,

	Conveyor,
	ExpressConveyor,
	Pusher,
	Gear,
	Laser,
}

export interface Tile {
	type : number
	dir ?: Directions
	num ?: number
	walls : Walls
}

export interface Board {
	tiles: Tile[][]
}
