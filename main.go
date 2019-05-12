package main

import (
	"github.com/tjbearse/robo/events"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/loader"
	"github.com/tjbearse/robo/websockets"
	"github.com/tjbearse/robo/bridge"
)

func main () {
	server := websockets.NewServer("")

	initialPhase := events.AddPlayerPhase{}
	board, err := loader.GetDefaultBoard()
	if err != nil {
		panic(err)
	}

	deck := loader.GetDefaultDeck()
	g := game.NewGame(*board, *deck, initialPhase)
	bridge := bridge.NewBridge(&g)

	server.Serve(&bridge)
}
