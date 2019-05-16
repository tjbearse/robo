export interface Card {
	Priority: number
	Command: Command
	Reps: number
}
export interface CardBack {
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

export type CardUpOrDown = Card | CardBack
