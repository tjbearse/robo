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
	RevealCard = 'NotifyRevealCard',
	PlayerReady = 'NotifyPlayerReady',
}

export default Notifications

