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

const phaseReducer = createReducer(Phases.Join, {
    [notify.StartSpawn]: (state, action) => Phases.Spawn
	// PlayerFinished = 'NotifyPlayerFinished',
})

/*
  UI related
  PromptForSpawn = 'PromptForSpawn',
*/

// root
import { combineReducers } from 'redux'

const rootReducer = combineReducers({
	players: playersReducer,
	board: boardReducer,
	phase: phaseReducer,
})

export default rootReducer
