import React from 'react'
import ReactDOM from 'react-dom'

import { Player } from '../types/player'
import { Card as CardType } from '../types/card'
import { Card, CardSlot} from './card'
import range from './range'

interface HandAndBoardProps {
	player: Player,
	selectedCard: number,
	selectedBoard: number,
	selectCard: (i?:number)=>void,
	selectBoard: (i?:number)=>void,
}
export default function HandAndBoard({player, selectedCard, selectedBoard, selectCard, selectBoard}) {
	let {hand, board} = player
	return (
		<div id="playArea">
			<Hand {...{hand, selectedCard, select: selectCard}} />
			<Board {...{board, selectedBoard, select: selectBoard}} />
		</div>
	)
}

interface HandProps {
	hand: CardType[],
	selectedCard: number,
	select: (number)=>void,
	deselect: ()=>void,
}
export function Hand({hand, selectedCard, select}) {
	return (
		<div id="hand">Hand
			<div id="handArea" >
			{
				hand.map((card, i) => {
					let selected = i === selectedCard
					let selectCard = ()=>select(i)
					return (
						<Card key={i} {...{card, selected, select: selectCard}} />
					)
				})
			}
			</div>
		</div>
	)
}

interface BoardProps {
	board: { [i: number]: CardType },
	selectedBoard: number,
	select: (i?:number)=>void,
}
export function Board({board, selectedBoard, select}: BoardProps) {
	return (
		<div id="robotBoard">Board:
			<div id="robotBoardArea">
				{
					range(5).map((i) => {
						let card = board[i]
						let selected = i === selectedBoard
						let selectCard = ()=>select(i)
						return card?
							   <Card key={i} {...{card, selected, select: selectCard}} /> :
							   <CardSlot key={i} {...{selected, slot:i, select: selectCard}} />
					})
				}
			</div>
		</div>
	)
}
