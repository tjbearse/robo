package bridge

import (
	"errors"
	"encoding/json"
	"fmt"
	
	"github.com/tjbearse/robo/websockets"
	"github.com/tjbearse/robo/events/comm"
)

type commEntity *websockets.Client

// start bridge file
type Bridge struct {
	clientToContext map[commEntity]map[comm.ContextKey]comm.ContextValue
	contextToClient map[comm.ContextKey]map[comm.ContextValue]map[commEntity]bool
	eventMap map[string]func()comm.IncomingEvent
}

func NewBridge(eventMap map[string]func()comm.IncomingEvent) Bridge {
	return Bridge{
		map[commEntity]map[comm.ContextKey]comm.ContextValue{},
		map[comm.ContextKey]map[comm.ContextValue]map[commEntity]bool{},
		eventMap,
	}
}

func (b *Bridge) associate(c commEntity, k comm.ContextKey, v comm.ContextValue) {
	if b.clientToContext[c] == nil {
		b.clientToContext[c] = map[comm.ContextKey]comm.ContextValue{}
	}
	b.clientToContext[c][k] = v

	if b.contextToClient[k] == nil {
		b.contextToClient[k] = map[comm.ContextValue]map[commEntity]bool{}
	}
	if b.contextToClient[k][v] == nil {
		b.contextToClient[k][v] = map[commEntity]bool{}
	}
	b.contextToClient[k][v][c] = true
}

func (b *Bridge) clear(c commEntity) {
	if b.clientToContext[c] == nil {
		return
	}
	for k, v := range(b.clientToContext[c]) {
		delete(b.contextToClient[k][v], c)
	}
	delete(b.clientToContext, c)
}

func (b *Bridge) getClients(k comm.ContextKey, v comm.ContextValue) []commEntity {
	m := b.contextToClient[k][v]
	if m == nil {
		return []commEntity{}
	}

    clients := make([]commEntity, len(m))
	i := 0
    for k := range m {
        clients[i] = k
		i++
    }
	return clients
}

func (b *Bridge) getContext(c commEntity, k comm.ContextKey) comm.ContextValue {
	return b.clientToContext[c][k]
}

// TODO remove direct websockets references
// e.g. unpack envelopes. *Client may be a little harder
func (c *Bridge) ActOnMessage(renv websockets.Envelope, send func(websockets.Envelope)) {
	comm := commClient{c, send, renv.Client}
	event, err := c.unpackIncomingEvent(renv.Msg);
	if err != nil {
		comm.SendError(err)
		return
	}
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in while acting on a message", r)
				comm.SendError(errors.New("there was an unexpected error. Your game might be broken. Please let TJ know!"));
			}
		}()
		err = event.Exec(comm)
	}();
	if err != nil {
		comm.SendError(err)
	}
}

type Envelope struct {
	Type string
	Msg interface{}
}



// decide which context we have before this point
func (bridge *Bridge) unpackIncomingEvent(b []byte) (comm.IncomingEvent, error) {
	// read envelope
	var msg json.RawMessage
	env := Envelope {
		Msg: &msg,
	}
	if err := json.Unmarshal(b, &env); err != nil {
		return nil, errors.New("Couldn't unpack message")
	}

	t, ok := bridge.eventMap[env.Type]
	if ok != true {
		return nil, errors.New("Event didn't match known type")
	}
	event := t()

	if err := json.Unmarshal(msg, event); err != nil {
		return nil, fmt.Errorf("Couldn't unpack message as %s, %v", env.Type, string(msg))
	}
	return event, nil
}
