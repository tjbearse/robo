// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websockets

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gobuffalo/packr"
)

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

var box packr.Box
var startTime = time.Now()

func init() {
	box = packr.NewBox("./templates")
}

func (s *server) Serve(b Bridge) {
	hub := newHub(b)
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fbytes, _ := box.Find("home.html")
	h := bytes.NewReader(fbytes)
	http.ServeContent(w, r, "home.html", startTime, h)
}
