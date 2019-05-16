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
	Forward = 0,
	Backward = 1,
	TurnLeft = 2,
	TurnRight = 3,
	UTurn = 4
}

export type CardUpOrDown = Card | CardBack
