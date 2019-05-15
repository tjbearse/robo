package events

import (
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/events/comm"
	"github.com/tjbearse/robo/game/loader"
)

// TODO players and games should reference a different type of store
type GameId int
var gameStore = map[GameId]*game.Game{}

type NewGame struct {
	PlayerName string
	// TODO GameName string
}

func (e NewGame) Exec(cc comm.CommClient) error {
	c,err := comm.WithoutContext(cc)

	// TODO better id schema
	id := GameId(len(gameStore))
	board, err := loader.GetDefaultBoard()
	if err != nil {
		return err
	}
	deck := loader.GetDefaultDeck()
	g := game.NewGame(
		*board,
		*deck,
		&AddPlayerPhase{},
	)
	gameStore[id] = &g
	c.Clear() // TODO also update client to clear state?
	c.SetGame(gameStore[id])
	c.Reply(NotifyNewGame{id})

	return JoinGame{e.PlayerName, id}.Exec(cc)
}

type GetGames struct {}

func (e GetGames) Exec(cc comm.CommClient) error {
	// TODO
	return nil
}
