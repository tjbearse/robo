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
	over := runRegisters(c,g)
	if over {
		return
	}
	StartSpawnPhase(cc)
}

type SimulationPhase struct {}

func runRegisters(c comm.ExtendedCommClient, g *game.Game) (gameOver bool) {
	for reg:=0; reg < Steps; reg++ {
		over := runRegister(c, g, reg)
		if over {
			return true
		}
	}
	cleanup(c, g)
	return false
}

type commandPair struct {
	Player *game.Player
	Card cards.Card
}

func runRegister(c comm.ExtendedCommClient, g *game.Game, reg int) (gameOver bool) {
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
	return touchCheckpoints(c, g)
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

// TODO pushersPush
func pushersPush(c comm.ExtendedCommClient, g *game.Game, reg int) {
}

func gearsRotate(c comm.ExtendedCommClient, g *game.Game) {
	players := g.GetPlayers()
	for p, _ := range(players) {
		if p.Robot.Configuration == nil {
			continue
		}
		r := p.Robot
		loc := r.Configuration.Location
		tile, err := g.Board.GetTile(loc)
		if err != nil {
			continue
		}
		if tile.Type == game.Gear {
			target := *r.Configuration
			if tile.Dir == coords.North {
				target.Heading = target.Heading.RotateRight(1)
			} else {
				target.Heading = target.Heading.RotateRight(-1)
			}
			c.Broadcast(g, NotifyRobotMoved{p.Name, Gear, *r.Configuration, target})
			*r.Configuration = target
		}
	}
}

func touchCheckpoints(c comm.ExtendedCommClient, g *game.Game) (gameOver bool) {
	players := g.GetPlayers()
	for p, _ := range(players) {
		if p.Robot.Configuration == nil {
			continue
		}
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
				return true
			}
		}
	}
	return false
}

func cleanup(c comm.ExtendedCommClient, g *game.Game) {
	// Repairs & Upgrades
	// TODO upgrades
	players := g.GetPlayers()
	for p, _ := range(players) {
		if p.Robot.Configuration == nil {
			continue
		}
		loc := p.Robot.Configuration.Location
		tile, err := g.Board.GetTile(loc)
		if err != nil {
			// TODO debug logging
			continue
		}
		switch tile.Type {
			case game.Repair:
			if p.Robot.Damage != 0 {
				p.Robot.Damage--
				c.Broadcast(g, NotifyHeal{p.Name, p.Robot.Damage})
			}
			case game.Upgrade:
			if p.Robot.Damage != 0 {
				p.Robot.Damage--
				if p.Robot.Damage != 0 {
					p.Robot.Damage--
				}
				c.Broadcast(g, NotifyHeal{p.Name, p.Robot.Damage})
			}
			case game.Flag:
		}
	}

	// Wiping Registers
	for p := range(players) {
		unlockedNum := HandSize - p.Robot.Damage
		if unlockedNum > Steps {
			unlockedNum = Steps
		}
		for i:= 0; i < unlockedNum; i++ {
			card := p.Robot.Board[i] 
			if card != nil {
				g.Deck.Discard(*card)
				p.Robot.Board[i] = nil
			}
		}
		c.Broadcast(g, NotifyCleanup{p.Name, p.Robot.Board})
	}
}

// TODO need to also message loss of a life
func executeRobotFall(p *game.Player, target coords.Configuration, reason MoveReason) NotifyRobotFell {
	r := &p.Robot
	r.Lives--
	e := NotifyRobotFell{p.Name, reason, *r.Configuration, target}
	r.Configuration = nil
	return e
}
