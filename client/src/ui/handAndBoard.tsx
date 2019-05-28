export default function HandAndBoard({player, selCard, selBoard}) { // TODO
	return null
}

/*
function drawMyHandAndBoard(myPlayer, cardSel, boardSel){
	let ePlayArea = tag('div')
	ePlayArea.id = 'playArea'

	// hand
	let eHand = tag('ol')
	eHand.id = 'hand'
	eHand.start = 0
	myPlayer.hand.forEach((card, i) => {
		let eCard = getCard(card)
		if (cardSel == i) {
			eCard.className += ' selected'
		} else {
			eCard.onclick = ()=> { SelectCard(i) }
		}
		eHand.appendChild(eCard)
	})
	let heading = tag('div')
	heading.innerText = 'Hand'
	heading.appendChild(eHand)
	ePlayArea.appendChild(heading)

	// robot board
	let eBoard = tag('ol')
	eBoard.id = 'robot-board'
	eBoard.start = 0
	for (let i=0; i < 5; i++) {
		let eSlot;
		if (myPlayer.board[i]) {
			eSlot = getCard(myPlayer.board[i])
		} else {
			eSlot = tag('li')
			eSlot.innerText = '____'
		}
		if (boardSel == i) {
			eSlot.className += ' selected'
		} else {
			eSlot.onclick = ()=> { SelectSlot(i) }
		}
		eBoard.appendChild(eSlot)
	}
	heading = tag('div')
	heading.innerText = 'Board'
	heading.appendChild(eBoard)
	ePlayArea.appendChild(heading)

	return ePlayArea

	function getCard(c : Card) {
		let eCard = tag('li')
		eCard.className = 'card'
		let text = commandToText(c.Command) + ' '
		if (c.Command === Command.Move) {
			text += `${c.Reps} `
		}
		eCard.innerText = text
		let ePrior = tag('span')
		ePrior.className = 'priority'
		ePrior.innerText = `(${c.Priority})`
		eCard.appendChild(ePrior)
		return eCard
	}
}
*/
