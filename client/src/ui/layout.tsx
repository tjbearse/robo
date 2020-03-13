import React from 'react'
import ReactDOM from 'react-dom'

let style = {
	width: '100vw',
	height: '100vh',
	display: 'block',
}

export default function Layout() {
	return (
		<canvas style={style} >
			<div> game board </div>
			<div> cards & my board </div>
			<div> player displays </div>
		</canvas>
	)
}
