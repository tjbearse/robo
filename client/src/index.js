import store from "./store"
// import './game.css'
import {conn, init} from './websocket'
import * as uiActions from './uiActions'
// import drawCrappyVersion from './ui'
import queue from './queue'
import './ui/app'

// This file is mostly temporary stuff while I work on data store

console.log(store.getState())
const unsubscribe = store.subscribe(() => console.log(store.getState()))

init()
conn.onmessage = handleMessage
conn.onclose = handleClose
// crappy player api :)
Object.assign(window, {
	...uiActions
})

function handleMessage(evt) {
	let messages = evt.data.split('\n');
	for (let i = 0; i < messages.length; i++) {
		let json = JSON.parse(messages[i])
		let type = json.Type
		let payload = json.Msg
		queue.push({type, payload})
	}
}

function handleClose() {
	store.dispatch({type: 'error', payload: 'connection closed'})
}
