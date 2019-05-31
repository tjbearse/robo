import { createReducer } from 'redux-starter-kit'
import notify from '../actions/notify'
import {ClearError} from '../actions/playerActions'
import {cardToBoard} from '../actions/playerTriggered'
import {Card, fromObject} from '../types/card'

/*
  UI related
  PromptForSpawn = 'PromptForSpawn',
*/

const initialState = {
	colors: {map: {}, count: 0},
	winner: '',
	currentMove: {
		playerName: '',
		card: null,
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


	[notify.RevealCard]: (state, { payload: {Player, Card} }) => {
		state.currentMove = { playerName: Player, card: fromObject(Card), }
	},
})

export default uiInfoReducer
