import {conn} from '../websocket'
import {Dir} from '../types/coord'
import store from '../store'

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

function SetSpawnHeading(dir : Dir) {
	let e = {
		Dir: dir
	}
	return sendEvent('SetSpawnHeading', e)
}

function NewGame(PlayerName : string) {
	let e = {
		PlayerName
	}
	return sendEvent('NewGame', e)
}

function GetGames() {
	return sendEvent('GetGames', {})
}

function CardToBoard(HandOffset:number, BoardSlot:number) {
	let e = {
		HandOffset,
		BoardSlot,
	}
	return sendEvent('CardToBoard', e)
}

function CardToHand(BoardSlot:number) {
	let e = {
		BoardSlot,
	}
	return sendEvent('CardToHand', e)
}

function CommitCards() {
	return sendEvent('CommitCards', {})
}

function JoinGame(PlayerName:string, Game:number) {
	let e = {
		PlayerName,
		Game,
	}
	return sendEvent('JoinGame', e)
}

function LeaveGame() {
	return sendEvent('LeaveGame', {})
}

function ReadyToSpawn() {
	return sendEvent('ReadyToSpawn', {})
}

function sendEvent(Type, Msg) {
	conn.send(JSON.stringify({ Type, Msg }))
	store.dispatch({
		type: Type,
		payload: Msg
	})
}
