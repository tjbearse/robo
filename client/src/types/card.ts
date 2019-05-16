class Card {
	Priority: number
	Command: Command
	Reps: number
}
class CardBack {
}

enum Command {
	Forward = 0,
	Backward = 1,
	TurnLeft = 2,
	TurnRight = 3,
	UTurn = 4
}

type CardUpOrDown = Card | CardBack

export {CardUpOrDown, CardBack, Card, Command}
