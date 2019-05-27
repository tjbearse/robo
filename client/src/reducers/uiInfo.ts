import { createReducer } from 'redux-starter-kit'
import notify from '../actions/notify'
import {ClearError} from '../actions/playerActions'

/*
  UI related
  PromptForSpawn = 'PromptForSpawn',
*/

const initialState = {
	colors: {map: {}, count: 0},
	winner: ''
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
})

export default uiInfoReducer
