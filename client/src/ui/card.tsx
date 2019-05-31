import React from 'react'
import ReactDOM from 'react-dom'

import { Card as CardType, Command, commandToText } from '../types/card'

interface CardProps {
	card: CardType,
	selected: boolean,
	select?: ()=>void,
}
export function Card(props: CardProps) {
	let {card, selected, select} = props
	if (card.Command === undefined) {
		return <CardBack {...props} />
	}
	let onClick = ()=>{ select && select() }
	let className = `card command-${Command[card.Command]}`
	if (selected) {
		className += ' selected'
	}
	let priority = (card.Priority||0).toString().padStart(3, '0')
	return (
		<div {...{className, onClick}} >
			<div className="moveImg" />
			<div className="priority">{ priority }</div>
			<div className="command">{commandToText(card.Command)}&nbsp;
				<span className="reps">{card.Reps}</span>
			</div>
		</div>
	)
}


export function CardSlot({selected, select, slot}) {
	let onClick = select
	let className = 'card empty'
	if (selected) {
		className += ' selected'
	}
	return (
		<div {...{className, onClick}}>
			<span>Slot {slot}</span>
		</div>
	)
}

export function CardBack({card, selected, select}: CardProps) {
	let onClick = ()=>{ select && select() }
	let className = 'card card-back'
	if (selected) {
		className += ' selected'
	}
	return (
		<div {...{className, onClick}} ></div>
	)
}
