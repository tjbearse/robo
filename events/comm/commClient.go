package comm

import (
	"errors"
	
	"github.com/tjbearse/robo/game"
)

type CommClient interface {
	Reply(OutgoingEvent)
	Message(OutgoingEvent, ContextKey, ContextValue)

	Associate(ContextKey, ContextValue)
	Clear()
	GetContext(ContextKey) ContextValue
}
type IncomingEvent interface {
	Exec(CommClient) error
}
type OutgoingEvent interface {}
type ContextKey interface{}
type ContextValue interface{}

const (
	gameKey int = iota
	playerKey int = iota
)

var noGameError = errors.New("Not associated to a game")
var noPlayerError = errors.New("Not associated with a player")

func WithoutContext(c CommClient) (ExtendedCommClient, error) {
	return ExtendedCommClient{c}, nil
}

func WithGameContext(c CommClient) (ExtendedCommClient, *game.Game, error) {
	e := ExtendedCommClient{c}
	g, _, err := e.unpack()
	if err != nil {
		return e, g, err
	}
	if g == nil {
		return e, g, noPlayerError
	}
	return e, g, nil
}

func WithPlayerContext(c CommClient) (ExtendedCommClient, *game.Game, *game.Player, error) {
	e := ExtendedCommClient{c}
	g, p, err := e.unpack()
	if err != nil {
		return e, g, p, err
	}
	if g == nil {
		return e, g, p, noPlayerError
	}
	if p == nil {
		return e, g, p, noPlayerError
	}
	return e, g, p, nil
}

type ExtendedCommClient struct {
	CommClient CommClient
}

func (e *ExtendedCommClient) unpack() (*game.Game, *game.Player, error) {
	gI := e.CommClient.GetContext(gameKey)
	g, ok := gI.(*game.Game)
	if !ok {
		return nil, nil, noGameError
	}

	pI := e.CommClient.GetContext(playerKey)
	p, ok := pI.(*game.Player)
	if !ok {
		return nil, nil, noPlayerError
	}

	return g, p, nil
}

func (e *ExtendedCommClient) Reply(o OutgoingEvent) {
	e.CommClient.Reply(o)
}

func (e *ExtendedCommClient) Broadcast(g *game.Game, o OutgoingEvent) {
	e.CommClient.Message(o, gameKey, g)
}

func (e *ExtendedCommClient) MessagePlayer(p *game.Player, o OutgoingEvent) {
	e.CommClient.Message(o, playerKey, p)
}

func (e *ExtendedCommClient) SetPlayer(p *game.Player) {
	e.CommClient.Associate(playerKey, p)
}

func (e *ExtendedCommClient) SetGame(g *game.Game) {
	e.CommClient.Associate(gameKey, g)
}

// TODO should be closing too?
func (e *ExtendedCommClient) Clear() {
	e.CommClient.Clear()
}
