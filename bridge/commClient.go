package bridge

import (
	"encoding/json"
	"reflect"
	
	"github.com/tjbearse/robo/events"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/websockets"
)

type commClient struct {
	bridge *Bridge
	send func(websockets.Envelope)
	client *websockets.Client
	player *game.Player
}

func (c *commClient) Broadcast(raw events.OutGoingEvent) {
	c.sendTo(raw, nil)
}

func (c *commClient) Message(raw events.OutGoingEvent, p *game.Player) {
	c.bridge.getClient(p)
	c.sendTo(raw, c.client)
}

func (c *commClient) sendTo(raw events.OutGoingEvent, cl *websockets.Client) {
	env := Envelope{
		Msg: raw,
		Type: reflect.TypeOf(raw).String(),
	}
	msg, _ := json.Marshal(env)
	c.send(websockets.Envelope{cl, msg})
}

func (c *commClient) Associate(p *game.Player) {
	c.bridge.associate(c.client, p)
}
func (c *commClient) Deassociate() {
	c.bridge.deassociate(c.client)
}
func (c *commClient) SendError(err error) {
	env := Envelope {
		"error",
		err.Error(),
	}
	msg, _ := json.Marshal(env)
	c.send(websockets.Envelope{c.client, msg})
}

func (c *commClient) GetPlayer() *game.Player {
	return c.player
}
