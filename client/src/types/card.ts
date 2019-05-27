export interface Card {
	Priority: number
	Command: Command
	Reps: number
}
export interface CardBack {
}

interface CardSerial {
	Priority: number
	Command: string
	Reps: number
}

export function fromObject(c : CardSerial) : Card {
	return {
		Priority: c.Priority,
		Command: Command[c.Command],
		Reps: c.Reps
	}
}
export function newCard(Priority: number, Command : Command, Reps: number) : Card {
	return {
		Priority,
		Command,
		Reps
	}
}
export function newCardBack() : CardBack {
	return {}
}

export enum Command {
	Move = 0,
	BackUp,
	RotateLeft,
	RotateRight,
	UTurn,
}

export function commandToText(c :Command) {
	switch(c) {
		case Command.Move:
			return 'Move'
		case Command.BackUp:
			return 'Back Up'
		case Command.RotateLeft:
			return 'Rotate Left'
		case Command.RotateRight:
			return 'Rotate Right'
		case Command.UTurn:
			return 'U-Turn'
		default:
			return 'Move'
	}
}

export type CardUpOrDown = Card | CardBack
