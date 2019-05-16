import store from "./store"
import {Walls, TileType} from "./types/board"
import './game.css'
import {conn, init} from './websocket'
import * as uiActions from './uiActions'

// This file is mostly temporary stuff while I work on data store

console.log(store.getState())
const unsubscribe = store.subscribe(() => console.log(store.getState()))

window.onload = windowOnLoad
// crappy player api :)
Object.assign(window, {
	...uiActions
})

// --

function windowOnLoad() {
	init()
	conn.onmessage = handleMessage

	drawCrappyVersion(store.getState())
	const unsubscribeDraw = store.subscribe(() => drawCrappyVersion(store.getState()))

	document.getElementById("form").onsubmit = submitForm
};

function handleMessage(evt) {
	let messages = evt.data.split('\n');
	for (let i = 0; i < messages.length; i++) {
		let json = JSON.parse(messages[i])
		let type = json.Type
		let payload = json.Msg
		store.dispatch({type, payload})
	}
}
function drawCrappyVersion(state) {
	const {
		players: {me, players},
		board, 
		phase
	} = state
	const myPlayer = players[me]

	let e = document.createElement('div')
	e.id = 'gameArea'
	
	e.appendChild(drawCrappyBoard(board, players))
	if (myPlayer) {
		e.appendChild(drawMyHandAndBoard(myPlayer))
	}

	let old = document.getElementById('gameArea')
	old.parentNode.replaceChild(e, old);
	return

	// --

	function drawCrappyBoard(board, players){
		let eBoard = document.createElement('table')
		eBoard.id = 'board'
		if (board.length == 0) {
			return eBoard
		}
		let coords = Object.values(players).reduce(( acc, player ) => {
			if (player.robot.config) {
				acc.push(player)
			}
			return acc
		}, []).reduce(( acc, player ) => {
			let robot = player.robot
			let {X, Y} = robot.config.Location
			if (!acc[X]) {
				acc[X] = {}
			}
			acc[X][Y] = player
			return acc
		}, {})
		/*
		  x -->
		  y
		  |
		  v
		*/
		for (let y=-1; y<board[0].length+1; y++) {
			let row = document.createElement('tr')
			for (let x=-1; x<board.length+1; x++) {
				let cell = document.createElement('td')
				if (x < 0 || x >= board.length || y < 0 || y >= board[0].length) {
					// Off map
					cell.className = 'tile tile-OffMap'
				} else {
					// Real map
					let tile = board[x][y]
					let wallClass = [Walls.North, Walls.East, Walls.South, Walls.West]
						.reduce((acc, w) => {
							if (tile.walls & w) {
								acc += 'wall-' + Walls[w] + ' '
							}
							return acc
						}, '')
					cell.className = 'tile tile-' + TileType[tile.type] +
						' dir-' + tile.dir
					if (wallClass) {
						let wall = document.createElement('div')
						wall.className = 'wall ' + wallClass
						cell.appendChild(wall)
					}
				}
				if (coords[x] && coords[x][y]) {
					let player = coords[x][y]
					let eRobot = document.createElement('div')
					eRobot.className = 'robot dir-'+(player.robot.config.Heading || 'indeterminent')
					// TODO differentiate players
					cell.appendChild(eRobot)
				}
				row.appendChild(cell)
			}
			eBoard.appendChild(row)
		}

		return eBoard
	}

	function drawMyHandAndBoard(myPlayer){
		let ePlayArea = document.createElement('div')
		ePlayArea.id = 'playArea'

		// hand
		let eHand = document.createElement('ol')
		eHand.id = 'hand'
		eHand.start = '0'
		myPlayer.hand.forEach((card) => {
			let eCard = getCard(card)
			eHand.appendChild(eCard)
		})
		let heading = document.createElement('div')
		heading.innerText = 'Hand'
		heading.appendChild(eHand)
		ePlayArea.appendChild(heading)

		// robot board
		let eBoard = document.createElement('ol')
		eBoard.id = 'robot-board'
		for (let i=0; i < 5; i++) {
			let eSlot;
			if (myPlayer.board[i]) {
				eSlot = getCard(myPlayer.board[i])
			} else {
				eSlot = document.createElement('li')
			}
			eBoard.appendChild(eSlot)
		}
		heading = document.createElement('div')
		heading.innerText = 'Board'
		heading.appendChild(eBoard)
		ePlayArea.appendChild(heading)

		return ePlayArea

		function getCard(c) {
			let eCard = document.createElement('li')
			eCard.className = 'card'
			eCard.innerText = JSON.stringify(Object.values(c))
			return eCard
		}
	}

}

function submitForm() {
	var msg = document.getElementById("Msg");
	var type = document.getElementById("Type");

	if (!conn) {
		return false;
	}
	if (!msg.value) {
		return false;
	}
	if (!type.value) {
		return false;
	}
	try {
		var msgV = JSON.parse(msg.value)
		var envelope = {
			Type: type.value,
			Msg: msgV
		};
		conn.send(JSON.stringify(envelope));
		msg.value = "";
		type.value = "";
	} catch {
		return false;
	}
	return false;
};

