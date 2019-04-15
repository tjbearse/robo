package events

import (
	"errors"

	"github.com/tjbearse/robo/game"
)

type SpawnPhase struct {
	waiting map[*game.Player]bool
}

// In:
// Spawn those without choices
// Prompt those with choices

// State:
// spawns TODO

// Actions:
// set spawn selection

// GOTO
// selectActions when all spawned

func NewSpawnPhase(g *game.Game) SpawnPhase {
	// who is not spawned
	assign := make([]*game.Player, 0)
	sel := map[*game.Player]bool{}
	players := g.GetPlayers()
	for p := range(players) {
		if p.Robot.Configuration == nil {
			if p.Spawn == nil {
				assign = append(assign, p)
			} else if p.Robot.Lives != 0 {
				sel[p] = false
			}
		}
	}
	for _, p := range(assign) {
		// FIXME using _
		c, _ := g.GetNextSpawn()
		ccopy := c
		p.Spawn = &c
		p.Robot.Configuration = &ccopy
		// TODO updated robot configuration
	}
	// for p := range(sel) {
		// TODO prompt for spawn orientation in p.Spawn
	// }
	return SpawnPhase{sel}
}

func SelectSpawnHeading(g *game.Game, p *game.Player, d game.Dir) Event {
	return func () error {
		uPhase := g.GetPhase()
		ph, ok := uPhase.(SpawnPhase)
		if !ok {
			return errors.New("Not the right phase")
		}

		// TODO swat if already true
		ph.waiting[p] = true
		c := p.Spawn
		c.Heading = d // FIXME use update system
		p.Robot.Configuration = c // FIXME use update system
		// TODO updated robot configuration

		for _,v := range(ph.waiting) {
			if v != true {
				return nil
			}
		}
		newPh := NewPlayCardsPhase(g)
		g.ChangePhase(newPh)
		return nil
	}
}
