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
	const myPlayer = players[me]

	return (
		<div>
			{ uiInfo.error? <Error clear={ClearError} error={uiInfo.error}/> : '' }
			{ uiInfo.winner && <GameOver winner={uiInfo.winner} /> }
			<div id="BoardAndControls">
				<div id="Controls" >
					<div> GameId: { gameInfo.id } </div>
					<HealthOverview {...{players}} />
					<Form/>
				</div>
				<Board {...{board, players, uiInfo}} />
			</div>
		</div>
	)
}

const mapStateToProps = (state /*, ownProps*/) => state
const mapDispatchToProps = {ClearError}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Layout)
