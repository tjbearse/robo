package bridge

import (
	"errors"
	"encoding/json"
	"fmt"
	"reflect"
	
	"github.com/tjbearse/robo/events"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/websockets"
)

const (
	AddPlayer string = "AddPlayer"
	AddSelection string = "AddSelection"
	Ready string = "Ready"
	RemovePlayer string = "RemovePlayer"
	SelectSpawnHeading string = "SelectSpawnHeading"
)

type commEntity interface{}

// start bridge file
type Bridge struct {
	game *game.Game
	cl2P map[commEntity]*game.Player
	p2Cl map[*game.Player]*websockets.Client
}

func NewBridge(g *game.Game) Bridge {
	return Bridge{
		g,
		map[commEntity]*game.Player{},
		map[*game.Player]*websockets.Client{},
	}
}

func (b *Bridge) associate(c *websockets.Client, p *game.Player) {
	b.cl2P[c] = p
	b.p2Cl[p] = c
}

func (b *Bridge) deassociate(c *websockets.Client) {
	p := b.cl2P[c]
	delete(b.cl2P, c)
	delete(b.p2Cl, p)
}

func (b *Bridge) getClient(p *game.Player) *websockets.Client {
	return b.p2Cl[p]
}


// TODO remove direct websockets references
// e.g. unpack envelopes. *Client may be a little harder
func (c *Bridge) ActOnMessage(renv websockets.Envelope, send func(websockets.Envelope)) {
	p, _ := c.cl2P[renv.Client];
	comm := commClient{c, send, renv.Client, p}
	event, err := unpackIncomingEvent(renv.Msg);
	if err != nil {
		comm.SendError(err)
		return
	}
	err = event.Exec(&comm, c.game)
	if err != nil {
		comm.SendError(err)
	}
}

type Envelope struct {
	Type string
	Msg interface{}
}



var eventMap = map[string]reflect.Type {
	"JoinGame": reflect.TypeOf(events.JoinGame{}),
	"LeaveGame": reflect.TypeOf(events.LeaveGame{}),
	"ReadyToSpawn": reflect.TypeOf(events.ReadyToSpawn{}),
	"SetSpawnHeading": reflect.TypeOf(events.SetSpawnHeading{}),
	"CardToBoard": reflect.TypeOf(events.CardToBoard{}),
	"CardToHand": reflect.TypeOf(events.CardToHand{}),
}

// decide which context we have before this point
func unpackIncomingEvent(b []byte) (events.IncomingEvent, error) {
	// read envelope
	var msg json.RawMessage
	env := Envelope {
		Msg: &msg,
	}
	if err := json.Unmarshal(b, &env); err != nil {
		return nil, errors.New("Couldn't unpack message")
		// TODO tell client to stop doing that
	}

	t, ok := eventMap[env.Type]
	if ok != true {
		return nil, errors.New("Event didn't match known type")
	}
	ghost := reflect.New(t).Interface()
	event, ok := ghost.(events.IncomingEvent)
	if !ok {
		return nil, errors.New("reflection unsuccessful")
	}

	if err := json.Unmarshal(msg, event); err != nil {
		return nil, fmt.Errorf("Couldn't unpack message as %s, %v", env.Type, string(msg))
	}
	return event, nil
}
