package events

import (
	"github.com/tjbearse/robo/events/comm"
	"github.com/tjbearse/robo/game"
	"github.com/tjbearse/robo/game/coords"
)

type ray struct {
	Start coords.Coord
	StartPlayer *game.Player // nil if env
	End coords.Coord
	EndPlayer *game.Player // nil if hits a wall
}
func fireLasers(c comm.ExtendedCommClient, g *game.Game, reg int) {
	// gather targets so hits are simultaneous
	rays := []ray{}
	for _, start := range(g.Board.GetLasers()) {
		tile, _ := g.Board.GetTile(start)
		end, p := rayTrace(g, start, tile.Dir)
		rays = append(rays, ray{start, nil, end, p})
	}
	for p := range(g.GetPlayers()) {
		if p.Robot.Configuration != nil {
			config := *p.Robot.Configuration
			end, p := rayTrace(g, config.Location, config.Heading)
			rays = append(rays, ray{config.Location, p, end, p})
		}
	}
	if len(rays) < 0 {
		return
	}

	referenceRays := make([]Ray, len(rays))
	for i, r := range(rays) {
		sName := ""
		if r.StartPlayer != nil {
			sName = r.StartPlayer.Name
		}
		eName := ""
		if r.EndPlayer != nil {
			eName = r.EndPlayer.Name
		}
		referenceRays[i] = Ray{r.Start, sName, r.End, eName}
	}
	c.Broadcast(g, NotifyLaserFire{referenceRays})
	for _, ray := range(rays) {
		if ray.EndPlayer != nil {
			r := ray.EndPlayer.Robot
			r.Damage++
			c.Broadcast(g, NotifyDamage{ray.EndPlayer.Name, r.Damage})
			if r.Damage == MaxDamage {
				r.Lives--
				r.Configuration = nil
				c.Broadcast(g, NotifyLifeLoss{ray.EndPlayer.Name, r.Lives})
			}
		}
	}
}

func rayTrace(g *game.Game, coord coords.Coord, dir coords.Dir) (coords.Coord, *game.Player) {
	if !g.Board.IsInbounds(coord) {
		return coord, nil
	}
	// no player check in the first square
	offset := coords.Offset{0,1}.Rotate(dir)
	newCoord := coord.Apply(offset)
	if ok, _ := g.Board.IsPassable(coord, newCoord); ok {
		return coord, nil
	}
	coord = newCoord
	for {
		if !g.Board.IsInbounds(coord) {
			return coord, nil
		}
		p := g.CheckForRobot(coord)
		if p != nil {
			return coord, p
		}

		offset := coords.Offset{0,1}.Rotate(dir)
		newCoord := coord.Apply(offset)
		if ok, _ := g.Board.IsPassable(coord, newCoord); ok {
			return coord, nil
		}
		coord = newCoord
	}
}
