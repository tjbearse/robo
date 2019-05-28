import React from 'react'
import ReactDOM from 'react-dom'

import { Player } from '../types/player'

interface PlayerMap {
	[name:string]: Player
}

export default function HealthOverview({ players } : { players: PlayerMap }) {
	let items = Object.values(players).map((player) => {
		let { name, hand, board, robot: { lives, damage } } = player
		let text = `${name}: lives: ${lives} damage: ${damage}`
		return (
			<li key="player-{name}">{text}</li>
		)
	})
	return (
		<div id="overview">
			<ol>
				{items}
			</ol>
		</div>
	)
}
