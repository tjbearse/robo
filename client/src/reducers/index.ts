import { createReducer } from 'redux-starter-kit'
import notify from '../actions/notify'
import playersReducer from './player'
import boardReducer from './board'

enum Phases {
	Join,
	Spawn,
	PlayCards,
	Simulate,
	GameOver,
}

const gameInfoReducer = createReducer({
	id: '',
	phase: Phases.Join
}, {
    [notify.StartSpawn]: (state, action) => { state.phase = Phases.Spawn },
	[notify.Welcome]: (state, { payload: { GameId }}) => ({ id: GameId, phase: Phases.Join })
	// PlayerFinished = 'NotifyPlayerFinished',
})

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
	[notify.Welcome]: (state, {Name}) => initialState,
})

// root
import { combineReducers } from 'redux'

const rootReducer = combineReducers({
	players: playersReducer,
	board: boardReducer,
	gameInfo: gameInfoReducer,
	uiInfo: uiInfoReducer,
})

export default rootReducer
