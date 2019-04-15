package events

// TODO this whole file
type RunPhase struct {
	Round int
}
	// In:
	// execute actions in order

	// Actions:
	// ready for next

	// GOTO
	// Game over when won
	// spawn when all ready
	// Simulate phase when all ready or timer

func (ph *RunPhase) runRound() {
	// TODO functionality needed in game to run these
	// run robots
	// run hazards
	// shoot robot lasers

	// TODO update status
	ph.Round++
	if ph.Round == Steps {
		// check victory
		// round wrap
	}
}
