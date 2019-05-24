enum Notifications {
	AddPlayer = 'NotifyAddPlayer',
	Welcome = 'NotifyWelcome',
	RemovePlayer = 'NotifyRemovePlayer',

	FlagTouched = 'NotifyFlagTouched',
	RobotFell = 'NotifyRobotFell',
	RobotMoved = 'NotifyRobotMoved',

	SpawnUpdate = 'NotifySpawnUpdate',
	PromptForSpawn = 'PromptForSpawn',

	// phase related
	StartSpawn = 'NotifyStartSpawn',
	PlayerFinished = 'NotifyPlayerFinished',

	// card related
	PromptWithHand = 'PromptWithHand',
	CardToBoard = 'NotifyCardToBoard',
	CardToBoardBlind = 'NotifyCardToBoardBlind',
	CardToHand = 'NotifyCardToHand',
	CardToHandBlind = 'NotifyCardToHandBlind',
	RandomBoardFill = 'NotifyRandomBoardFill',
	RevealCard = 'NotifyRevealCard',
	PlayerReady = 'NotifyPlayerReady',
	Cleanup = 'NotifyCleanup',

	// uncat
	Board = 'NotifyBoard',
	NewGame = 'NotifyNewGame',
	Error = "error",
	ErrorReport = "ErrorReport",
}

export default Notifications

