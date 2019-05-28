import { connect } from 'react-redux'
import React from 'react'
import ReactDOM from 'react-dom'

import { ClearError } from '../actions/playerActions'

import Form from './form'
import HealthOverview from './HealthOverview'
import Error from './error'
import Board from './board'
import HandAndBoard from './handAndBoard'
import GameOver from './gameOver'


function Layout (state) {
	const {
		players: {me, players},
		board, 
		phase,
		uiInfo,
		gameInfo,
	} = state
	const selCard = uiInfo.selected.card,
		  selBoard = uiInfo.selected.board
	const myPlayer = players[me]

	// TODO implement these components
	return (
		<div>
			<div> GameId: { gameInfo.id } </div>
			{ uiInfo.error? <Error clear={ClearError} error={uiInfo.error}/> : '' }
			<HealthOverview {...{players}} />
			<Board {...{board, players, uiInfo}} />
			<Form/>
			{ myPlayer && <HandAndBoard player={myPlayer} selCard={selCard} selBoard={selBoard} /> }
			{ uiInfo.winner && <GameOver winner={uiInfo.winner} /> }
		</div>
	)
}

const mapStateToProps = (state /*, ownProps*/) => state
const mapDispatchToProps = {ClearError}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Layout)
