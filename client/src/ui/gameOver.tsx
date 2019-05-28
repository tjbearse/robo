import React from 'react'
import ReactDOM from 'react-dom'

export default function GameOver({winner}) {
	return (
		<span id="gameOver">{winner} won!</span>
	)
}
