import React from 'react'
import ReactDOM from 'react-dom'

import { Player } from '../types/player'

interface HealthOverviewProps {
	players: {[name:string]: Player},
	colorMap: {[name:string]: number},
}
export default function HealthOverview({ players, colorMap } : HealthOverviewProps) {
	let items = Object.values(players).map((player) => {
		let { name, hand, board, robot: { lives, damage } } = player
		return (
			<li key="player-{name}">
				{`${name}:`}
				<span className={`robot robot-${colorMap[name]}`} />
				{`lives: ${lives} damage: ${damage}`}
			</li>
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
