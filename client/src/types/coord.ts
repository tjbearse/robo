export interface Coord {
	X: number
	Y: number
}

export interface Config {
	Location: Coord
	Heading: Dir
}

export enum Dir {
	Indeterminent = "",
	North = "North",
	East = "East",
	South = "South",
	West = "West",
}
