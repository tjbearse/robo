package main

import (
	"github.com/tjbearse/robo/events"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/websockets"
	"github.com/tjbearse/robo/bridge"
)

func main () {
	server := websockets.NewServer("")

	initialPhase := events.AddPlayerPhase{}
	g := game.NewGame(game.Board{}, game.NewDeck(), initialPhase)
	bridge := bridge.NewBridge(&g)

	server.Serve(&bridge)
}
