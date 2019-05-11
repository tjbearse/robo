package events

import (
	"github.com/tjbearse/robo/game"
)

// Reasons for movement
type MoveReason string
const (
	Spawned MoveReason = "spawned"
	Moved = "moved"
)
type NotifyRobotMoved struct {
	Name string
	Reason MoveReason
	OldConfig game.Configuration
	NewConfig game.Configuration
}
