package main

import (
	"net/http"

	"github.com/tjbearse/robo/bridge"
	"github.com/tjbearse/robo/events"
	"github.com/tjbearse/robo/websockets"
)


func main () {
	var fs http.FileSystem = http.Dir("client/dist")

	bridge := bridge.NewBridge(events.EventMap)

	server := websockets.NewServer("", &bridge, fs)
	server.Serve()
}
