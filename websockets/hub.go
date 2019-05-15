// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websockets

import (
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	inbound chan Envelope

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	bridge Bridge
}

// make this a pump instead?
type Bridge interface {
	ActOnMessage(Envelope, func(Envelope))
}

type Envelope struct {
	Client *Client
	Msg []byte
}


// buffer this?
func newHub(b Bridge) *Hub {
	return &Hub{
		inbound:  make(chan Envelope),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		bridge: b,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.inbound:
			fmt.Printf("in: %s\n", message.Msg)
			h.bridge.ActOnMessage(message, h.NonBlockingSend)
		}
	}
}

func (h *Hub) NonBlockingSend(e Envelope) {
	fmt.Printf("out: %s\n", e.Msg)
	if e.Client == nil {
		// TODO obsolete, don't need catch all sends
		for client := range h.clients {
			select {
			case client.send <- e.Msg:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	} else {
		// FIXME: bridge doesn't realize clients get closed and keeps associations to them
		if h.clients[e.Client] {
			select {
			case e.Client.send <- e.Msg:
			default:
				close(e.Client.send)
				delete(h.clients, e.Client)
			}
		}
	}
}
