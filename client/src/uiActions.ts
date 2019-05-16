import {conn} from './websocket'
import {Dir} from './types/coord'
import {
	CardToBoard,
	CardToHand,
	CommitCards,
	GetGames,
	JoinGame,
	LeaveGame,
	NewGame,
	ReadyToSpawn,
	SetSpawnHeading,
} from './actions/playerTriggered'

export function newGame() {
	let playerName = getFormString('name')
	return NewGame(playerName)
}

export function joinGame() {
	let playerName = getFormString('name')
	let gameId = getFormInt('gameId')
	return JoinGame(playerName, gameId)
}

export function readyToSpawn() {
	ReadyToSpawn()
}

export function cardToBoard() {
	let HandOffset = getFormInt('cardslot')
	let BoardSlot = getFormInt('boardslot')
	CardToBoard(HandOffset, BoardSlot)
}

export function cardToHand() {
	let BoardSlot = getFormInt('boardslot')
	CardToHand(BoardSlot)
}

export function commitCards() {
	CommitCards()
}

export function setSpawnHeading() {
	var e = document.getElementById("spawnHeading") as HTMLSelectElement
	var heading = e.options[e.selectedIndex].text as Dir
	SetSpawnHeading(heading)
}

export function getGames() {
	GetGames()
}

export function leaveGame() {
	LeaveGame()
}


function getFormString(id:string) : string {
	let elm : HTMLInputElement = document.getElementById(id) as HTMLInputElement
	return elm.value
}

function getFormInt(id:string) : number {
	let elm : HTMLInputElement = document.getElementById(id) as HTMLInputElement
	return +elm.value
}
