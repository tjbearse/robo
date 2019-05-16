package main

import (
	"net/http"
	"flag"

	"github.com/tjbearse/robo/bridge"
	"github.com/tjbearse/robo/events"
	"github.com/tjbearse/robo/websockets"
)


func main () {
	addr := flag.String("addr", ":8081", "help message for flagname")
	flag.Parse()

	var fs http.FileSystem = http.Dir("client/dist")

	bridge := bridge.NewBridge(events.EventMap)

	server := websockets.NewServer(*addr, &bridge, fs)
	server.Serve()
}
