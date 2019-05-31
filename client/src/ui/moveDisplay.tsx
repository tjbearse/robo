import React from 'react'
import ReactDOM from 'react-dom'
import { connect } from 'react-redux'
import { Card } from './card'

function MoveDisplay({player: { name, color}, card}) {
	if (name === '') {
		return null
	}

	return (
		<div id="MoveDisplay">
			<div>
				<div className={`robot robot-${color}`}></div>
				<div className="name">{name}</div>
			</div>
			<Card card={card} />
		</div>
	)
}

const mapStateToProps = ({uiInfo: { currentMove, colors} }) => ({
	player: {
		name: currentMove.playerName,
		color: colors.map[currentMove.playerName],
	},
	card: currentMove.card,
})
export default connect(
  mapStateToProps,
)(MoveDisplay)
