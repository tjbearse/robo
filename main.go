package main

import "github.com/tjbearse/robo/websockets"

func main () {
	server := websockets.NewServer("")
	server.Serve()
}
