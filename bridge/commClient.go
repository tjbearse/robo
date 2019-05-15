package bridge

import (
	"encoding/json"
	"reflect"
	"strings"
	
	"github.com/tjbearse/robo/websockets" // TODO remove
	"github.com/tjbearse/robo/events/comm"
)

type commClient struct {
	bridge *Bridge
	send func(websockets.Envelope)
	client commEntity
}

func (c commClient) Reply(raw comm.OutgoingEvent) {
	c.sendTo(raw, c.client)
}

func (c commClient) Message(raw comm.OutgoingEvent, k comm.ContextKey, v comm.ContextValue) {
	clients := c.bridge.getClients(k,v)
	for _, client := range(clients) {
		c.sendTo(raw, client)
	}
}

func (c commClient) sendTo(raw comm.OutgoingEvent, cl commEntity) {
	// FIXME this is hacky, do this properly
	t := reflect.TypeOf(raw).String()
	parts := strings.Split(t, ".")
	t = parts[len(parts) -1]
	env := Envelope{
		Msg: raw,
		Type: t,
	}
	msg, _ := json.Marshal(env)
	c.send(websockets.Envelope{cl, msg})
}

func (c commClient) Associate(key comm.ContextKey, val comm.ContextValue) {
	c.bridge.associate(c.client, key, val)
}

func (c commClient) Clear() {
	c.bridge.clear(c.client)
}

func (c commClient) SendError(err error) {
	env := Envelope {
		"error",
		err.Error(),
	}
	msg, _ := json.Marshal(env)
	c.send(websockets.Envelope{c.client, msg})
}

func (c commClient) GetContext(key comm.ContextKey) comm.ContextValue {
	return c.bridge.getContext(c.client, key)
}
