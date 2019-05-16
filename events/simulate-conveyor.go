package events

import (
	"github.com/tjbearse/robo/events/comm"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/coords"
)

func moveConveyors(c comm.ExtendedCommClient, g *game.Game) {
	conveyorPhases := [] func(game.TileType) bool {
		func(t game.TileType) bool { return t==game.ExpressConveyor },
		func(t game.TileType) bool { return t==game.ExpressConveyor || t==game.Conveyor },
	}
	for _, conveyorSelector := range(conveyorPhases) {
		moves := getMoves(g, conveyorSelector)
		eliminateImpassible(g, moves)
		eliminateShared(moves)
		changed := eliminateStaticCollisions(g, moves)
		for changed {
			changed = eliminateStaticCollisions(g, moves)
		}

		// execute moves
		for p, coord := range(moves) {
			oldConfig := *p.Robot.Configuration
			heading := determineHeading(g, oldConfig.Heading, oldConfig.Location, coord)
			destination := coords.Configuration{coord, heading}
			c.Broadcast(g, NotifyRobotMoved{p.Name, Conveyed, oldConfig, destination})
			*p.Robot.Configuration = destination
		}
	}
}

func getMoves(g *game.Game, moving func(game.TileType)bool) (map[*game.Player]coords.Coord) {
	moves := map[*game.Player]coords.Coord{}
	for p := range(g.GetPlayers()) {
		r := &p.Robot
		if (r.Configuration != nil) {
			coord := r.Configuration.Location
			tile, err := g.Board.GetTile(coord)
			if err != nil {
				continue
			}
			if moving(tile.Type) {
				moves[p] = getConveyCoord(coord, tile)
			}
		}
	}
	return moves
}


func eliminateImpassible(g *game.Game, moves map[*game.Player]coords.Coord) {
	for p, target := range(moves) {
		if !g.Board.IsInbounds(target) {
			continue
		}
		curr := p.Robot.Configuration.Location
		passable, _ := g.Board.IsPassable(curr, target) 
		if !passable {
			delete(moves, p)
		}
	}
}

func eliminateShared(moves map[*game.Player]coords.Coord) {
	for p, target := range(moves) {
		for p2, target2 := range(moves) {
			if p != p2 && target == target2 {
				delete(moves, p)
				delete(moves, p2)
			}
		}
	}
}

func eliminateStaticCollisions(g *game.Game, moves map[*game.Player]coords.Coord) (changed bool) {
	changed = false
	for p, target := range(moves) {
		collidePlayer := g.CheckForRobot(target)
		if collidePlayer != nil {
			if _, ok := moves[collidePlayer]; !ok {
				changed = true
				delete(moves, p)
			}
		}
	}
	return
}

func getConveyCoord(coord coords.Coord, tile game.Tile) (coords.Coord) {
	offset := coords.Offset{0,1}.Rotate(tile.Dir)
	return coord.Apply(offset)
} 


func determineHeading(g *game.Game, currHeading coords.Dir, currPos, nextPos coords.Coord) coords.Dir {
	newTile, _ := g.Board.GetTile(nextPos)
	if newTile.Type == game.ExpressConveyor || newTile.Type == game.Conveyor {
		oldTile, _ := g.Board.GetTile(currPos)
		new := newTile.Dir
		old := oldTile.Dir
		// not same, nor opposites
		if new != old && new.RotateRight(2) != old {
			if old.RotateRight(1) == new {
				// right turn
				return currHeading.RotateRight(1)
			} else {
				// left turn
				return currHeading.RotateRight(3)
			}
		}
	}
	return currHeading
}
