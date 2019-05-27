import { createReducer } from 'redux-starter-kit'
import notify from '../actions/notify'
import {SelectCard, SelectSlot, ClearError} from '../actions/playerActions'
import {CardToBoard} from '../actions/playerTriggered'

/*
  UI related
  PromptForSpawn = 'PromptForSpawn',
*/

const initialState = {
	colors: {map: {}, count: 0},
	winner: '',
	selected: {
		board: 0,
		card: 0,
	},
}
const uiInfoReducer = createReducer(initialState, {
	[notify.AddPlayer]: (state, {payload: {Name}}) => {
		state.colors.map[Name] = state.colors.count
		state.colors.count++
	},
	[notify.PlayerFinished]: (state, {payload: {Player}}) => {
		state.winner = Player
	},
	[notify.Welcome]: (state, action) => initialState,
	[notify.Goodbye]: (state, action) => initialState,
	[notify.ErrorReport]: (state, {payload: {Error}})=> { state.error = Error },
	[notify.Error]: (state, {payload}) => { state.error = payload },
	[ClearError.type]: (state, action) => { state.error = '' },

	[notify.PromptWithHand]: (state, action) => { state.selected = initialState.selected },
	[SelectCard.type]: (state, action) => { state.selected.card = action.payload },
	[SelectSlot.type]: (state, action) => { state.selected.board = action.payload },
	[CardToBoard.type]: (state,action) => { state.selected.board++ },
})

export default uiInfoReducer
