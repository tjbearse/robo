import { createReducer } from 'redux-starter-kit'
import notify from '../actions/notify'
import playersReducer from './player'

//
const boardReducer = createReducer([], {
})

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
/*
// Types
	tile: {
		Type
		Dir
		Num (flags and initial spawns)
		Walls: N,E,S,W
	}

// Store
	board: {
		tiles: [][]
	},
*/
