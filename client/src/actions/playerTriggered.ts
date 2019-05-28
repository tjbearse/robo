import {conn} from '../websocket'
import {Dir} from '../types/coord'
import store from '../store'

setSpawnHeading.type = 'SetSpawnHeading'
newGame.type = 'NewGame'
getGames.type = 'GetGames'
cardToBoard.type = 'CardToBoard'
cardToHand.type = 'CardToHand'
commitCards.type = 'CommitCards'
joinGame.type = 'JoinGame'
leaveGame.type = 'LeaveGame'
readyToSpawn.type = 'ReadyToSpawn'

interface action {
	type: string
	payload: any
}

export {
	cardToBoard,
	cardToHand,
	commitCards,
	getGames,
	joinGame,
	leaveGame,
	newGame,
	readyToSpawn,
	setSpawnHeading,
}

function setSpawnHeading(dir : Dir) {
	let e = {
		Dir: dir
	}
	return sendEvent('SetSpawnHeading', e)
}

function newGame(PlayerName : string) {
	let e = {
		PlayerName
	}
	return sendEvent('NewGame', e)
}

function getGames() {
	return sendEvent('GetGames', {})
}

function cardToBoard(HandOffset:number, BoardSlot:number) {
	let e = {
		HandOffset,
		BoardSlot,
	}
	return sendEvent('CardToBoard', e)
}

function cardToHand(BoardSlot:number) {
	let e = {
		BoardSlot,
	}
	return sendEvent('CardToHand', e)
}

function commitCards() {
	return sendEvent('CommitCards', {})
}

function joinGame(PlayerName:string, Game:number) {
	let e = {
		PlayerName,
		Game,
	}
	return sendEvent('JoinGame', e)
}

function leaveGame() {
	return sendEvent('LeaveGame', {})
}

function readyToSpawn() {
	return sendEvent('ReadyToSpawn', {})
}

function sendEvent(Type, Msg) {
	conn.send(JSON.stringify({ Type, Msg }))
	return {
		type: Type,
		payload: Msg
	}
}
