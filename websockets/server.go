// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websockets

import (
	"log"
	"net/http"
	"time"

)


type server struct {
	Addr string
	Bridge Bridge
	FS http.FileSystem
}

func NewServer(addr string, bridge Bridge, fs http.FileSystem) (server) {
	if (addr == "") {
		addr = ":8080"
	}
	return server{
		addr,
		bridge,
		fs,
	}
}

var startTime = time.Now()

func init() {
}

func (s *server) Serve() {
	hub := newHub(s.Bridge)
	go hub.run()

	http.Handle("/", http.FileServer(s.FS))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
