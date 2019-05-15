package main

import (
	"github.com/tjbearse/robo/events"
	"github.com/tjbearse/robo/websockets"
	"github.com/tjbearse/robo/bridge"
)

func main () {
	server := websockets.NewServer("")

	bridge := bridge.NewBridge(events.EventMap)

	server.Serve(&bridge)
}
