import { createReducer } from 'redux-starter-kit'

import notify from '../actions/notify'
import {Player, newPlayer} from '../types/player'

/*
  me: string(name) | null
  players
*/
const initialState = { me: null, players: {} }
// TODO add welcome message to golang so we know who "me" is
const playersReducer = createReducer(
	initialState,
	stripActionFluff({
		[notify.AddPlayer]: ({players}, {Name}) => { players[Name] = newPlayer(Name) },
		[notify.RemovePlayer]: ({players}, {Name}) => { delete players[Name] },
		[notify.Welcome]: (state, {Name}) => { state.me = Name },

		...onMe({
			[notify.PromptWithHand]: (me, {Cards}) => { me.hand = Cards },
			[notify.CardToBoard]: (me, payload) => {
				const {BoardSlot, HandOffset} = payload
				me.board[BoardSlot] = me.hand.splice(HandOffset, 1)[0]
			},
			[notify.CardToHand]: (me, payload) => {
				const {BoardSlot, HandOffset} = payload
				me.hand.splice(HandOffset, 0, me.board[BoardSlot])
				delete me.board[BoardSlot]
			},
		}),
		...onPlayer({
			[notify.SpawnUpdate]: (p:Player, {Coord}) => { p.spawn = Coord },

			[notify.RobotMoved]: (p:Player, {NewConfig}) => { p.robot = NewConfig },

			/*
			  FlagTouched = 'NotifyFlagTouched',
			  RobotFell = 'NotifyRobotFell',
			  RobotMoved = 'NotifyRobotMoved',

			  PlayerReady = 'NotifyPlayerReady',

			  //not me
			  CardToBoardBlind = 'NotifyCardToBoardBlind',
			  CardToHandBlind = 'NotifyCardToHandBlind',
			  RevealCard = 'NotifyRevealCard',
			*/
		}),
	})
)
export default playersReducer

// --

function stripActionFluff(reducers: {[s: string]: (state, payload) => any}) {
	return mapObject(reducers, ([k,v]) => [k, strip(v)])

	function strip(reducer) {
		return function(state, {payload}) {
			return reducer(state, payload)
		}
	}
}

function onMe(reducers: {[s: string]: (p: Player, payload) => any} ) {
	return mapObject(reducers, ([k,v]) => [k, selectMe(v)])

	function selectMe(fn: (p:Player, payload) => any) {
		return function(state, payload) {
			let {me, players} = state
			const mePlayer = players[me]
			fn(mePlayer, payload) // must modify in place
		}
	}
}
function onPlayer(reducers: {[s: string]: (Player, payload) => any} ) {
	return mapObject(reducers, ([k,v]) => [k, selectToPlayer(v)])

	function selectToPlayer(fn: (p:Player, payload) => any) {
		return function(state, payload) {
			let {me, players} = state
			const name = payload.Name
			const player = players[name]
			fn(player, payload) // must modify in place
		}
	}

}

function mapObject(o, fn) {
	return fromEntries(Object.entries(o).map(fn))
}
function fromEntries(o :[any, any][]) {
	return o.reduce((acc, [k,v]) => { acc[k] = v; return acc }, {})
}