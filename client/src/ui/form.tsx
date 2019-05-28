import React from 'react'
import ReactDOM from 'react-dom'
import { connect } from 'react-redux'

import {Dir} from '../types/coord'
import Phases from '../types/phases'
import * as uiActions from '../actions/playerTriggered'

function Form(props ) {
	let {
		gameInfo,
		uiInfo,

		cardToBoard,
		cardToHand,
		commitCards,
		joinGame,
		leaveGame,
		newGame,
		readyToSpawn,
		setSpawnHeading,
	} = props
	let card = uiInfo.selected.card
	let board = uiInfo.selected.board
	let inside = getInnerContent()
	return (
		<div id="ControlForm">
			{ gameInfo.phase != Phases.NoGame && <button onClick={leaveGame}>Leave Game</button>}
			{ inside }
		</div>
	)

	// -- 

	function getInnerContent() {
		switch(gameInfo.phase) {
			case Phases.NoGame:
				return (<JoinGame {...{newGame, joinGame}} />)
			case Phases.Join:
				return (<button onClick={readyToSpawn}>Ready To Spawn</button>)
			case Phases.SpawnWait:
				return (<div>'Waiting for others to spawn'</div>)
			case Phases.Spawn:
				return (<SetSpawnHeading {...{setSpawnHeading}}/>)
			case Phases.PlayCards:
				return (<PlayCards {...{card, board, cardToBoard, cardToHand, commitCards}}/>)
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

function PlayCards ({card, board, cardToBoard, cardToHand, commitCards}) {
	let cToB = () => cardToBoard(card,board)
	let cToH = () => cardToHand(board)
	let commit = ()=> commitCards()
	return (
		<div>
			<button onClick={cToB}>Card To Board</button>
			<button onClick={cToH}>Card To Hand</button>
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

const mapStateToProps = (state /*, ownProps*/) => state
const mapDispatchToProps = uiActions
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Form)
