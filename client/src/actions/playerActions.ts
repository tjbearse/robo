import store from '../store'

function ClearError() {
	store.dispatch({
		type: 'ClearError',
	})
}
ClearError.type = 'ClearError'

function SelectCard(card : number) {
	store.dispatch({
		type: 'SelectCard',
		payload: card,
	})
}
SelectCard.type = 'SelectCard'

function SelectSlot(slot : number) {
	store.dispatch({
		type: 'SelectSlot',
		payload: slot,
	})
}
SelectSlot.type = 'SelectSlot'

export {
	ClearError,
	SelectSlot,
	SelectCard,
}
