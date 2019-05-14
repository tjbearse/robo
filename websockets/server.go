// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websockets

import (
	"log"
	"net/http"
	"time"

	"github.com/gobuffalo/packr"
)

const usePackr = false

type server struct {
	Addr string
}

func NewServer(addr string) (server) {
	if (addr == "") {
		addr = ":8080"
	}
	return server{
		addr,
	}
}

var fs http.FileSystem
var startTime = time.Now()

func init() {
	if usePackr {
		fs = packr.NewBox("../client/dist")
	} else {
		fs = http.Dir("./client/dist")
	}
}

func (s *server) Serve(b Bridge) {
	hub := newHub(b)
	go hub.run()

	http.Handle("/", http.FileServer(fs))

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
