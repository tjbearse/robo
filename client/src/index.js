import store from "./store"
import './game.css'
import {conn, init} from './websocket'
import * as uiActions from './uiActions'
import drawCrappyVersion from './ui'

// This file is mostly temporary stuff while I work on data store

console.log(store.getState())
const unsubscribe = store.subscribe(() => console.log(store.getState()))

window.onload = windowOnLoad
// crappy player api :)
Object.assign(window, {
	...uiActions
})

// --

function windowOnLoad() {
	init()
	conn.onmessage = handleMessage
	conn.onclose = handleClose

	drawCrappyVersion(store.getState())
	const unsubscribeDraw = store.subscribe(() => drawCrappyVersion(store.getState()))

	document.getElementById("form").onsubmit = submitForm
};

function handleMessage(evt) {
	let messages = evt.data.split('\n');
	for (let i = 0; i < messages.length; i++) {
		let json = JSON.parse(messages[i])
		let type = json.Type
		let payload = json.Msg
		store.dispatch({type, payload})
	}
}

function handleClose() {
	store.dispatch({type: 'error', payload: 'connection closed'})
}

function submitForm() {
	var msg = document.getElementById("Msg");
	var type = document.getElementById("Type");

	if (!conn) {
		return false;
	}
	if (!msg.value) {
		return false;
	}
	if (!type.value) {
		return false;
	}
	try {
		var msgV = JSON.parse(msg.value)
		var envelope = {
			Type: type.value,
			Msg: msgV
		};
		conn.send(JSON.stringify(envelope));
		msg.value = "";
		type.value = "";
	} catch {
		return false;
	}
	return false;
};

