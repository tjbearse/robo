package events

import (
	"sort"

	"github.com/tjbearse/robo/events/comm"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/cards"
	"github.com/tjbearse/robo/game/coords"
)

func StartSimulationPhase(cc comm.CommClient) {
	c, g, err := comm.WithGameContext(cc)
	if err != nil {
		return // TODO
	}
	g.ChangePhase(&SimulationPhase{})
	runTurn(c,g)
	StartSpawnPhase(cc)
}

type SimulationPhase struct {}

func runTurn(c comm.ExtendedCommClient, g *game.Game) {
	for reg:=0; reg < Steps; reg++ {
		runRegister(c, g, reg)
	}
	cleanup(c, g)
}

type commandPair struct {
	Player *game.Player
	Card cards.Card
}

func runRegister(c comm.ExtendedCommClient, g *game.Game, reg int) {
	// Flip Cards
	commands := flipCards(g, reg)
	for _, command := range(commands) {
		c.Broadcast(g, NotifyRevealCard{command.Player.Name, command.Card})
	}

	// Move Robots
	for _, command := range(commands) {
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
	sort.Slice(commands, func (i, j int) bool { return cards.CardCompare(commands[i].Card, commands[j].Card) })
	return commands
}

func resolveMove(c comm.ExtendedCommClient, g *game.Game, cp commandPair) {
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
				c.Broadcast(g, NotifyRobotMoved{p.Name, Moved, previous, desiredMove})
				*p.Robot.Configuration = desiredMove
			}
		}
		attemptMove(c, g, desiredMove.Location, p, Moved)
		if p.Robot.Configuration == nil {
			return
		}
	}
}

// TODO should these events be collected together? Perhaps return those instead
func attemptMove(c comm.ExtendedCommClient, g *game.Game, loc coords.Coord, p *game.Player, reason MoveReason) (success bool) {
	r := &p.Robot
	if r.Configuration == nil {
		return false
	}
	target := coords.Configuration{loc, r.Configuration.Heading}
	if !g.Board.IsInbounds(loc) {
		c.Broadcast(g, executeRobotFall(p, target, reason))
		return true
	}
	// FIXME ignored error
	passable, _ := g.Board.IsPassable(r.Configuration.Location, loc) 
	if !passable {
		return false
	}

	tile, _ := g.Board.GetTile(loc)
	if tile.Type == game.Pit {
		c.Broadcast(g, executeRobotFall(p, target, reason))
		return true
	}
	collidePlayer := g.CheckForRobot(loc)
	if collidePlayer != nil {
		prev := r.Configuration.Location
		// Note this only works as long as offset steps are small (Len 1)
		collisionOffset := prev.OffsetTo(loc)
		collisionLoc := loc.Apply(collisionOffset)
		possible := attemptMove(c, g, collisionLoc, collidePlayer, Bumped)
		if !possible {
			return false
		}
	}
	// moveThere
	c.Broadcast(g, NotifyRobotMoved{p.Name, reason, *r.Configuration, target})
	*r.Configuration = target
	return true
}

func moveConveyors(c comm.ExtendedCommClient, g *game.Game) {
	// move Express Conveyors
	// move all conveyors
}

// TODO pushersPush
func pushersPush(c comm.ExtendedCommClient, g *game.Game, reg int) {
}

// TODO gearsRotate
func gearsRotate(c comm.ExtendedCommClient, g *game.Game) {
}

// TODO lasers fire
func fireLasers(c comm.ExtendedCommClient, g *game.Game, reg int) {
}

func touchCheckpoints(c comm.ExtendedCommClient, g *game.Game) {
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
				c.Broadcast(g, NotifySpawnUpdate{p.Name, loc})
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
				c.Broadcast(g, NotifyFlagTouched{p.Name, p.FlagNum})
				p.FlagNum++

				if p.FlagNum == g.Board.GetNumFlags() {
					StartGameWon(c.CommClient, p)
				}
			}
		}
	}
}

func cleanup(c comm.ExtendedCommClient, g *game.Game) {
	// Repairs & Upgrades
	// Wiping Registers
}

func executeRobotFall(p *game.Player, target coords.Configuration, reason MoveReason) NotifyRobotFell {
	r := &p.Robot
	r.Lives--
	e := NotifyRobotFell{p.Name, reason, *r.Configuration, target}
	r.Configuration = nil
	return e
}
