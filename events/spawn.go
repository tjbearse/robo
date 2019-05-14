package events

import (
	"errors"

	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/coords"
)

type SpawnPhase struct {
	waiting map[*game.Player]bool
}

func StartSpawnPhase(c commClient, g *game.Game) {
	c.Broadcast(NotifyStartSpawn{})

	selecting := map[*game.Player]bool{}
	for p := range(g.GetPlayers()) {
		// who is not on the board?
		if p.Robot.Configuration == nil {
			// FIXME all spawns after the first get 2 damage
			switch p.Spawn.State {
			case game.Unset:
				config, err := g.GetNextSpawn()
				if err != nil {
					// TODO tell everyone that everything is bad
				}
				p.Spawn.State = game.Rotatable
				p.Spawn.Coord = config.Location
				c.Broadcast(NotifySpawnUpdate{p.Name, config.Location})
				spawn(c, p, config)
			case game.Rotatable:
				loc := p.Spawn.Coord
				prompt := PromptForSpawn{p.Name, loc}
				c.Message(prompt, p)
				selecting[p] = false
			}
		}
	}

	if len(selecting) == 0 {
		StartCardsPhase(c, g)
	} else {
		g.ChangePhase(&SpawnPhase{selecting})
	}
}

type SetSpawnHeading struct {
	Dir coords.Dir
}
func (e SetSpawnHeading) Exec(c commClient, g *game.Game) error {
	p, err := getPlayer(c)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*SpawnPhase)
	if !ok {
		return errors.New("Not the right phase")
	}

	if _, ok := ph.waiting[p]; ok != true {
		return errors.New("We didn't ask for a spawn at this time")
	}
	delete(ph.waiting, p)
	config := coords.Configuration{p.Spawn.Coord, e.Dir}
	spawn(c, p, config)

	if len(ph.waiting) != 0 {
		return nil
	}

	StartCardsPhase(c, g)
	return nil
}

func spawn(comm commClient, p *game.Player, c coords.Configuration) {
	r := &p.Robot
	r.Configuration = &c
	comm.Broadcast(NotifyRobotMoved{p.Name, Spawned, coords.Configuration{}, c})
}
