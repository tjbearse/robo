import { createReducer } from 'redux-starter-kit'
import notify from '../actions/notify'
import * as outgoing from '../actions/playerTriggered'
import playersReducer from './player'
import boardReducer from './board'
import uiInfoReducer from './uiInfo'
import Phases from '../types/phases'

const initialState = {
	id: '',
	phase: Phases.NoGame
}
const gameInfoReducer = createReducer(initialState, {
	[notify.Welcome]: (state, { payload: { GameId }}) => ({ id: GameId, phase: Phases.Join }),
	[notify.Goodbye]: (state, action) => initialState,
    [notify.StartSpawn]: (state, action) => { state.phase = Phases.SpawnWait },
    [notify.PromptForSpawn]: (state, action) => { state.phase = Phases.Spawn },
    [notify.PromptWithHand]: (state, action) => { state.phase = Phases.PlayCards },
	[outgoing.commitCards.type]: (state, action) => { state.phase = Phases.PlayCardsWait },
	[notify.RevealCard]: (state, action) => { state.phase = Phases.Simulate },
	[notify.PlayerFinished]: (state, action) => { state.phase = Phases.GameOver },
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
