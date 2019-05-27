import {Walls, TileType, Tile} from "./types/board"
import {ClearError} from './actions/playerActions'
import {Player} from "./types/player"
import * as uiActions from './uiActions'
import Phases from './types/phases'
import { Card, Command, commandToText } from './types/card'
import { SelectCard, SelectSlot } from './actions/playerActions'

interface PlayerMap {
	[name: string]: Player
}

export default function drawCrappyVersion(state) {
	const {
		players: {me, players},
		board, 
		phase,
		uiInfo,
		gameInfo,
	} = state
	const myPlayer = players[me]

	let e = tag('div')
	e.id = 'gameArea'
	let eGameId = tag('div')
	eGameId.innerText = 'GameId: ' + gameInfo.id
	e.appendChild(eGameId)
	if (uiInfo.error) {
		e.appendChild(drawError(uiInfo.error))
	}
	e.appendChild(drawHealthOverview(players))
	
	e.appendChild(drawCrappyBoard(board, players, uiInfo))
	e.appendChild(drawCrappyForm(gameInfo, uiInfo))

	if (myPlayer) {
		e.appendChild(drawMyHandAndBoard(myPlayer, uiInfo.selected.card, uiInfo.selected.board))
	}
	if (uiInfo.winner) {
		e.appendChild(drawGameOver(uiInfo.winner))
	}


	let old = document.getElementById('gameArea')
	old.parentNode.replaceChild(e, old);
	return

	// --

	function drawHealthOverview(players : PlayerMap) {
		let eOverview = tag('div')
		eOverview.id = 'overview'
		eOverview.innerText = 'Players:'
		let list = tag('ol')
		let coords = Object.values(players).reduce(( e, player ) => {
			let { name, hand, board, robot: { lives, damage } } = player
			let elm = tag('li')
			elm.innerText = `${name}: lives: ${lives} damage: ${damage}`
			e.appendChild(elm)
			return e
		}, list)
		eOverview.appendChild(list)
		return eOverview
	}

	function drawGameOver(winner) {
		let span = tag('span')
		span.innerText = winner + ' won!'
		let e = tag('div')
		e.appendChild(span)
		e.id = 'gameOver'
		return e
	}

	function drawError(e) {
		let eError = tag('div')
		eError.id='error'
		eError.innerText = uiInfo.error
		eError.onclick = function() { ClearError() }
		return eError
	}

	function drawCrappyBoard(board, players : PlayerMap, { colors }){
		let eBoard = tag('table')
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
			let row = tag('tr')
			for (let x=-1; x<board.length+1; x++) {
				let cell = tag('td')
				if (x < 0 || x >= board.length || y < 0 || y >= board[0].length) {
					// Off map
					cell.className = 'tile tile-OffMap'
				} else {
					// Real map
					let tile : Tile = board[x][y]
					let wallClass = [Walls.North, Walls.East, Walls.South, Walls.West]
						.reduce((acc, w) => {
							if (tile.walls & w) {
								acc += 'wall-' + Walls[w] + ' '
							}
							return acc
						}, '')
					cell.className = 'tile tile-' + TileType[tile.type] +
						' dir-' + tile.dir
					if (tile.type == TileType.Flag) {
						cell.innerText += tile.num + 1
					}
					if (wallClass) {
						let wall = tag('div')
						wall.className = 'wall ' + wallClass
						cell.appendChild(wall)
					}
				}
				if (coords[x] && coords[x][y]) {
					let player = coords[x][y]
					let eRobot = tag('div')
					eRobot.className = 'robot dir-'+(player.robot.config.Heading || 'indeterminent') +
						' robot-' + colors.map[player.name]
					cell.appendChild(eRobot)
				}
				row.appendChild(cell)
			}
			eBoard.appendChild(row)
		}

		return eBoard
	}

	function drawMyHandAndBoard(myPlayer, cardSel, boardSel){
		let ePlayArea = tag('div')
		ePlayArea.id = 'playArea'

		// hand
		let eHand = tag('ol')
		eHand.id = 'hand'
		eHand.start = 0
		myPlayer.hand.forEach((card, i) => {
			let eCard = getCard(card)
			if (cardSel == i) {
				eCard.className += ' selected'
			} else {
				eCard.onclick = ()=> { SelectCard(i) }
			}
			eHand.appendChild(eCard)
		})
		let heading = tag('div')
		heading.innerText = 'Hand'
		heading.appendChild(eHand)
		ePlayArea.appendChild(heading)

		// robot board
		let eBoard = tag('ol')
		eBoard.id = 'robot-board'
		eBoard.start = 0
		for (let i=0; i < 5; i++) {
			let eSlot;
			if (myPlayer.board[i]) {
				eSlot = getCard(myPlayer.board[i])
			} else {
				eSlot = tag('li')
				eSlot.innerText = '____'
			}
			if (boardSel == i) {
				eSlot.className += ' selected'
			} else {
				eSlot.onclick = ()=> { SelectSlot(i) }
			}
			eBoard.appendChild(eSlot)
		}
		heading = tag('div')
		heading.innerText = 'Board'
		heading.appendChild(eBoard)
		ePlayArea.appendChild(heading)

		return ePlayArea

		function getCard(c : Card) {
			let eCard = tag('li')
			eCard.className = 'card'
			let text = commandToText(c.Command) + ' '
			if (c.Command === Command.Move) {
				text += `${c.Reps} `
			}
			eCard.innerText = text
			let ePrior = tag('span')
			ePrior.className = 'priority'
			ePrior.innerText = `(${c.Priority})`
			eCard.appendChild(ePrior)
			return eCard
		}
	}

}

function tag(s) {
	return document.createElement(s)
}
function button(str, onclickfn) {
	let btn = tag('button')
	btn.innerText = str
	btn.onclick = ()=> { onclickfn() }
	return btn
}

function drawCrappyForm(gameInfo, uiInfo) {
	let e = tag('div')
	e.id = 'ControlForm'
	if (gameInfo.phase != Phases.NoGame) {
		e.appendChild(button('Leave Game', uiActions.leaveGame))
	}
	switch(gameInfo.phase) {
		case Phases.NoGame:
			addJoinGame(e)
			break;
		case Phases.Join:
			let btn = button('Ready To Spawn', uiActions.readyToSpawn)
			e.appendChild(btn)
			break;
		case Phases.SpawnWait:
			addText(e, 'Waiting for others to spawn')
			break;
		case Phases.Spawn:
			addSetSpawnHeading(e)
			break;
		case Phases.PlayCards:
			addPlayCards(e, uiInfo.selected.card, uiInfo.selected.board)
			break;
		case Phases.PlayCardsWait:
			addText(e, 'Waiting for others to finish')
			break;
		case Phases.Simulate:
			addText(e, 'Running turns')
			break;
		case Phases.GameOver:
			addText(e, 'Game Over')
			break;
	}
	return e

	// --

	function addText(e, str) {
		let s = tag('span')
		s.innerText = str
		e.appendChild(s)
	}

	function addJoinGame(e) {
		let div = tag('div')
		div.innerHTML += `<label>name</label> <input value="TJ" id="name"/>
				<label>gameid</label> <input id="gameId"/>
				<br/>`
		let btn = tag('button')
		btn.innerText = 'New Game'
		btn.onclick= ()=> { uiActions.newGame() }
		div.appendChild(btn)
		btn = tag('button')
		btn.innerText = 'Join Game'
		btn.onclick= ()=> { uiActions.joinGame() }
		div.appendChild(btn)
		e.appendChild(div)
	}

	function addPlayCards(e, card, board) {
		let div = tag('div')
		// TODO remove this and rewire values into buttons below directly
		div.innerHTML = `<label>Hand Slot</label><input disabled id="cardslot" value="${card}" label="hand slot" type="number" />
		<label>Board Slot</label><input disabled id="boardslot" value="${board}" label="board slot" type="number" />
		<br/>`
		div.appendChild(button('Card To Board', uiActions.cardToBoard))
		div.appendChild(button('Card To Hand', uiActions.cardToHand))
		div.appendChild(button('Commit Cards', uiActions.commitCards))
		e.appendChild(div)
	}

	function addSetSpawnHeading(e) {
		let div = tag('div')
		div.innerHTML = `<select id="spawnHeading">
		  <option>North</option>
		  <option>East</option>
		  <option>South</option>
		  <option>West</option>
		</select>`
		div.appendChild(button('Set Spawn Heading', uiActions.setSpawnHeading))
		e.appendChild(div)
	}
}
