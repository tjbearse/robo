class Card {
	Priority: number
	Command: Command
	Reps: number
}

enum Command {
	Forward = 0,
	Backward = 1,
	TurnLeft = 2,
	TurnRight = 3,
	UTurn = 4
}

export {Card, Command}