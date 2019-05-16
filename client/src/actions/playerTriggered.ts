import {conn} from '../websocket'
import {Dir} from '../types/coord'

SetSpawnHeading.type = 'SetSpawnHeading'
NewGame.type = 'NewGame'
GetGames.type = 'GetGames'
CardToBoard.type = 'CardToBoard'
CardToHand.type = 'CardToHand'
CommitCards.type = 'CommitCards'
JoinGame.type = 'JoinGame'
LeaveGame.type = 'LeaveGame'
ReadyToSpawn.type = 'ReadyToSpawn'

interface action {
	type: string
	payload: any
}

export {
	CardToBoard,
	CardToHand,
	CommitCards,
	GetGames,
	JoinGame,
	LeaveGame,
	NewGame,
	ReadyToSpawn,
	SetSpawnHeading,
}

function SetSpawnHeading(dir : Dir) : action {
	let e = {
		Dir: dir
	}
	return sendEvent('SetSpawnHeading', e)
}

function NewGame(PlayerName : string) : action {
	let e = {
		PlayerName
	}
	return sendEvent('NewGame', e)
}

function GetGames() {
	return sendEvent('GetGames', {})
}

function CardToBoard(HandOffset:number, BoardSlot:number) : action {
	let e = {
		HandOffset,
		BoardSlot,
	}
	return sendEvent('CardToBoard', e)
}

function CardToHand(BoardSlot:number) : action {
	let e = {
		BoardSlot,
	}
	return sendEvent('CardToHand', e)
}

function CommitCards() : action {
	return sendEvent('CommitCards', {})
}

function JoinGame(PlayerName:string, Game:number) : action {
	let e = {
		PlayerName,
		Game,
	}
	return sendEvent('JoinGame', e)
}

function LeaveGame() :action {
	return sendEvent('LeaveGame', {})
}

function ReadyToSpawn() : action {
	return sendEvent('ReadyToSpawn', {})
}

function sendEvent(Type, Msg) : action {
	conn.send(JSON.stringify({ Type, Msg }))
	return {
		type: Type,
		payload: Msg
	}
}
