import React from 'react'
import ReactDOM from 'react-dom'
import { connect } from 'react-redux'

import {Dir} from '../types/coord'
import Phases from '../types/phases'
import { Player } from '../types/player'
import * as uiActions from '../actions/playerTriggered'
import HandAndBoard from './handAndBoard'


interface FormProps {
	phase: Phases,
	me: Player,

	cardToBoard,
	cardToHand,
	commitCards,
	joinGame,
	leaveGame,
	newGame,
	readyToSpawn,
	setSpawnHeading,
}
interface FormState {
	selectedCard: number,
	selectedBoard: number,
}

class Form extends React.Component {
	props: FormProps
	state: FormState
	constructor(props: FormProps) {
		super(props)
		this.state = {
			// TODO instead of unset, set board to the first empty slot
			selectedCard: null,
			selectedBoard: null,
		}
	}

	selectCard(i?: number) {
		let {selectedCard, selectedBoard} = this.state
		if (!this.props.me) {
			return
		}
		let hand = this.props.me.hand

		if (i === null) {
			selectedCard = null
		} else if (i < 0 || i >= hand.length) {
			console.error('selectCard called with out of index value', i)
			return
		} else if (i === selectedCard) {
			selectedCard = null
		} else if (selectedBoard != null) {
			this.props.cardToBoard(i, selectedBoard)
			selectedCard = null
			selectedBoard = null
		} else {
			selectedCard = i
		}
		this.setState({ selectedCard, selectedBoard })
	}

	selectBoard(i?: number) {
		let {selectedCard, selectedBoard} = this.state
		if (!this.props.me) {
			return
		}
		const board = this.props.me.board

		if (i == null) {
			selectedBoard = null
		/*
			// FIXME these length checks aren't easy on objects
		} else if (i < 0 || i >= board.length) {
			console.error('selectBoard called with out of index value', i)
			return
		*/
		} else if (board[i]) {
			this.props.cardToHand(i)
			selectedBoard = null
		} else if (i === selectedBoard) {
			selectedBoard = null
		} else if (selectedCard !== null) {
			this.props.cardToBoard(selectedCard, i)
			selectedBoard = null
			selectedCard = null
		} else {
			selectedBoard = i
		}
		this.setState({ selectedCard, selectedBoard })
	}

	render() {
		let {
			me,
			phase,

			cardToBoard,
			cardToHand,
			commitCards,
			joinGame,
			leaveGame,
			newGame,
			readyToSpawn,
			setSpawnHeading,
		} = this.props
		let { selectedCard, selectedBoard } = this.state
		let inside = getInnerContent()
		let player = me

		let selectCard = (i?: number)=>{ this.selectCard(i) }
		let selectBoard = (i?: number)=>{ this.selectBoard(i) }
		let showHand = player && phase != Phases.Join
		return (
			<div id="ControlForm">
				{ phase != Phases.NoGame &&
					<button onClick={leaveGame}>Leave Game</button>
				}
				{ inside }
				{ showHand && <HandAndBoard {...{player, selectedCard, selectedBoard, selectCard, selectBoard}} /> }
			</div>
		)

		// -- 

		function getInnerContent() {
			switch(phase) {
				case Phases.NoGame:
					return (<JoinGame {...{newGame, joinGame}} />)
				case Phases.Join:
					return (<button onClick={readyToSpawn}>Ready To Spawn</button>)
				case Phases.SpawnWait:
					return (<div>'Waiting for others to spawn'</div>)
				case Phases.Spawn:
					return (<SetSpawnHeading {...{setSpawnHeading}}/>)
				case Phases.PlayCards:
					return (<PlayCards {...{selectedCard, selectedBoard, cardToBoard, cardToHand, commitCards}}/>)
				case Phases.PlayCardsWait:
					return (<div>Waiting for others to finish</div>)
				case Phases.Simulate:
					return (<div>Running the turn</div>)
				case Phases.GameOver:
					return (<div>Game Over</div>)
			}
			return false
		}
	}
}

class HandleChange extends React.Component {
	constructor(props) {
		super(props)
		this.handleChange = this.handleChange.bind(this)
	}
	handleChange(event) {
		const target = event.target.name
		const value = event.target.value
		this.setState({[target]: value})
	}
}

interface JoinGameProps{
	newGame,
	joinGame,
}
interface JoinGameState{
	name: string,
	gameId: string|number,
}
class JoinGame extends HandleChange {
	props: JoinGameProps
	state: JoinGameState
	constructor(props: JoinGameProps) {
		super(props)
		this.state = {
			name: 'TJ',
			gameId: '',
		}
	}

	joinGame() {
		this.props.joinGame(this.state.name, this.state.gameId)
	}
	newGame() {
		this.props.newGame(this.state.name)
	}
	
	render() {
		return (
			<div>
				<label>name</label>
				<input onChange={this.handleChange} value={this.state.name} name="name"/>
				<label>gameid</label>
				<input onChange={this.handleChange} value={this.state.gameId} name="gameId"/>
				<br/>
				<button onClick={()=>this.newGame()}>New Game</button>
				<button onClick={()=>this.joinGame()}>Join Game</button>
			</div>
		)
	}
}

function PlayCards ({selectedCard, selectedBoard, cardToBoard, cardToHand, commitCards}) {
	let commit = ()=> commitCards()
	return (
		<div>
			<button onClick={commit}>Commit Cards</button>
		</div>
	)
}

interface SetSpawnHeadingProps {
	setSpawnHeading,
}
interface SetSpawnHeadingState {
	spawnHeading: Dir
}
class SetSpawnHeading extends HandleChange {
	props: SetSpawnHeadingProps
	state: SetSpawnHeadingState
	constructor(props: SetSpawnHeadingProps) {
		super(props)
		this.state = {
			spawnHeading: Dir.North,
		}
	}
	setSpawnHeading() {
		this.props.setSpawnHeading(this.state.spawnHeading)
	}
	render() {
		return (
			<div>
				<select name="spawnHeading" onChange={this.handleChange}>
					<option value={Dir.North}>North</option>
					<option value={Dir.East}>East</option>
					<option value={Dir.South}>South</option>
					<option value={Dir.West}>West</option>
				</select>
				<button onClick={()=>this.setSpawnHeading()}>Set Spawn Heading</button>
			</div>
		)
	}
}

const mapStateToProps = (state /*, ownProps*/) => ({
	phase: state.gameInfo.phase,
	me: state.players.players[state.players.me],
})
const mapDispatchToProps = uiActions
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Form)
