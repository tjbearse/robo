package events

import (
	"sort"

	"github.com/tjbearse/robo/game"
)

func StartSimulationPhase(c commClient, g *game.Game) {
	// notify simulation phase begun
	g.ChangePhase(&SimulationPhase{})
	runTurn(c,g)
	StartSpawnPhase(c, g)
}

type SimulationPhase struct {}

func runTurn(c commClient, g *game.Game) {
	for reg:=0; reg < Steps; reg++ {
		runRegister(c, g, reg)
	}
	cleanup(c, g)
}

type commandPair struct {
	Player *game.Player
	Card game.Card
}

func runRegister(c commClient, g *game.Game, reg int) {
	// Flip Cards
	commands := flipCards(g, reg)

	// Move Robots
	for _, command := range(commands) {
		c.Broadcast(NotifyRevealCard{command.Player.Name, command.Card})
		resolveMove(c, g, command)
	}

	moveConveyors(c, g)
	pushersPush(c, g, reg)
	gearsRotate(c, g)
	fireLasers(c, g, reg)
	touchCheckpoints(c, g)
}

func flipCards(g *game.Game, round int) []commandPair {
	players := g.GetPlayers()
	commands := make([]commandPair, 0, len(players))
	for p, _ := range(players) {
		if p.Robot.Lives > 0 && round < len(p.Robot.Board) && p.Robot.Board[round] != nil {
			card := *p.Robot.Board[round]
			commands = append(commands, commandPair{p, card})
		}
	}
	sort.Slice(commands, func (i, j int) bool { return game.CardCompare(commands[i].Card, commands[j].Card) })
	return commands
}

func resolveMove(c commClient, g *game.Game, cp commandPair) {
	p := cp.Player
	if p.Robot.Configuration == nil {
		return
	}
	previous := *p.Robot.Configuration
	steps := cp.Card.Apply(previous)

	for _, desiredMove := range(steps) {
		// Rotations always succeed
		if desiredMove.Location == previous.Location {
			if desiredMove.Heading != previous.Heading {
				c.Broadcast(NotifyRobotMoved{p.Robot.Name, Moved, previous, desiredMove})
				*p.Robot.Configuration = desiredMove
			}
		}
		attemptMove(c, g, desiredMove.Location, &p.Robot, Moved)
		if p.Robot.Configuration == nil {
			return
		}
	}
}

// TODO should these events be collected together? Perhaps return those instead
func attemptMove(c commClient, g *game.Game, loc game.Coord, r *game.Robot, reason MoveReason) (success bool) {
	if r.Configuration == nil {
		return false
	}
	target := game.Configuration{loc, r.Configuration.Heading}
	if !g.Board.IsInbounds(loc) {
		robotFalls(c, r, target, reason)
		return true
	}
	// FIXME ignored error
	passable, _ := g.Board.IsPassable(r.Configuration.Location, loc) 
	if !passable {
		return false
	}

	tile, _ := g.Board.GetTile(loc)
	if tile.Type == game.Pit {
		robotFalls(c, r, target, reason)
		return true
	}
	collideRobot := g.CheckForRobot(loc)
	if collideRobot != nil {
		prev := r.Configuration.Location
		// Note this only works as long as offset steps are small (Len 1)
		collisionOffset := prev.OffsetTo(loc)
		collisionLoc := loc.Apply(collisionOffset)
		possible := attemptMove(c, g, collisionLoc, collideRobot, Bumped)
		if !possible {
			return false
		}
	}
	// moveThere
	c.Broadcast(NotifyRobotMoved{r.Name, reason, *r.Configuration, target})
	*r.Configuration = target
	return true
}

func moveConveyors(c commClient, g *game.Game) {
	// move Express Conveyors
	// move all conveyors
}

// TODO pushersPush
func pushersPush(c commClient, g *game.Game, reg int) {
}

// TODO gearsRotate
func gearsRotate(c commClient, g *game.Game) {
}

// TODO lasers fire
func fireLasers(c commClient, g *game.Game, reg int) {
}

func touchCheckpoints(c commClient, g *game.Game) {
	players := g.GetPlayers()
	for p, _ := range(players) {
		if p.Robot.Configuration != nil {
			loc := p.Robot.Configuration.Location
			tile, err := g.Board.GetTile(loc)
			if err != nil {
				// TODO debug logging
				continue
			}
			// touching for spawn point
			if tile.Type == game.Repair ||
				tile.Type == game.Flag ||
				tile.Type == game.Upgrade {
				c.Broadcast(NotifySpawnUpdate{p.Name, loc})
				p.Spawn.State = game.Rotatable
				p.Spawn.Coord = loc
			}

			// touching next flag
			nextFlag, err := g.Board.GetFlag(p.FlagNum)
			if err != nil {
				// TODO debug logging
				continue
			}
			if tile.Type == game.Flag && loc == nextFlag {
				c.Broadcast(NotifyFlagTouched{p.Name, p.FlagNum})
				p.FlagNum++

				if p.FlagNum == g.Board.GetNumFlags() {
					StartGameWon(c, g, p)
				}
			}
		}
	}
}

func cleanup(c commClient, g *game.Game) {
	// Repairs & Upgrades
	// Wiping Registers
}

func robotFalls(c commClient, r *game.Robot, target game.Configuration, reason MoveReason) {
	r.Lives--
	c.Broadcast(NotifyRobotFell{r.Name, reason, *r.Configuration, target})
	r.Configuration = nil
}
