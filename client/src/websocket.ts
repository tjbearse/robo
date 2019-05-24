
var conn

function init() {
	if (!window["WebSocket"]) {
		window.alert("Your browser does not support WebSockets.")
	}
	let protocol = "ws://"
	if (document.location.protocol === 'https:') {
		protocol = "wss://"
	}
	conn = new WebSocket(protocol + document.location.host + document.location.pathname + "ws")
}

export { conn, init }
