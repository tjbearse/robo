package events

import (
	"errors"

	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/coords"
	"github.com/tjbearse/robo/events/comm"
)

type SpawnPhase struct {
	waiting map[*game.Player]bool
}

func StartSpawnPhase(cc comm.CommClient) {
	c, g, err := comm.WithGameContext(cc)
	if err != nil {
		// TODO err handle
	}
	c.Broadcast(g, NotifyStartSpawn{})

	selecting := map[*game.Player]bool{}
	for p := range(g.GetPlayers()) {
		// who is not on the board?
		if p.Robot.Configuration == nil {
			if p.Robot.Lives == 0 {
				continue
			}
			// FIXME don't spawn if dead
			switch p.Spawn.State {
			case game.Unset:
				config, err := g.GetNextSpawn()
				if err != nil {
					// TODO tell everyone that everything is bad
				}
				p.Spawn.State = game.Rotatable
				p.Spawn.Coord = config.Location
				c.Broadcast(g, NotifySpawnUpdate{p.Name, config.Location})
				c.Broadcast(g, executeSpawn(p, config))
			case game.Rotatable:
				loc := p.Spawn.Coord
				prompt := PromptForSpawn{p.Name, loc}
				// All spawns after the first get 2 damage
				p.Robot.Damage = 2
				c.Broadcast(g, NotifyDamage{p.Name, p.Robot.Damage})
				c.MessagePlayer(p, prompt)
				selecting[p] = false
			}
		}
	}

	if len(selecting) == 0 {
		StartCardsPhase(cc)
	} else {
		g.ChangePhase(&SpawnPhase{selecting})
	}
}

type SetSpawnHeading struct {
	Dir coords.Dir
}
func (e SetSpawnHeading) Exec(cc comm.CommClient) error {
	c, g, p, err := comm.WithPlayerContext(cc)
	if err != nil {
		return err
	}

	uPhase := g.GetPhase()
	ph, ok := uPhase.(*SpawnPhase)
	if !ok {
		return wrongPhaseError
	}

	if _, ok := ph.waiting[p]; ok != true {
		return errors.New("We didn't ask for a spawn at this time")
	}
	delete(ph.waiting, p)
	config := coords.Configuration{p.Spawn.Coord, e.Dir}
	c.Broadcast(g, executeSpawn(p, config))

	if len(ph.waiting) != 0 {
		return nil
	}

	StartCardsPhase(cc)
	return nil
}

func executeSpawn(p *game.Player, c coords.Configuration) NotifyRobotMoved {
	r := &p.Robot
	r.Configuration = &c
	return NotifyRobotMoved{p.Name, Spawned, coords.Configuration{}, c}
}
