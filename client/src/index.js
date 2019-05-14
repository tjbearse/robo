import store from "./store"

console.log(store.getState())
const unsubscribe = store.subscribe(() => console.log(store.getState()))
window.onload = windowOnLoad
var conn;
window.sendEvent = sendEvent

// --

function dispatchAction(wsFormat) {
	let action = {
		type: wsFormat.Type,
		payload: wsFormat.Msg
	}
	store.dispatch(action)
}

function sendEvent(Type, Msg) {
	conn.send(JSON.stringify({ Type, Msg }));
}

function windowOnLoad() {
    var msg = document.getElementById("Msg");
    var type = document.getElementById("Type");
    var log = document.getElementById("log");


    document.getElementById("form").onsubmit = function () {
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
			return false;
		} catch {
			return false;
		}
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
				var json = JSON.parse(messages[i])
				dispatchAction(json)
				var txt = JSON.stringify(json, null, 4)
                var item = document.createElement("pre");
                item.innerText = txt;
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
	return
	// --

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }
};

