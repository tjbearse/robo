package game


type GameStage int
// Stages of the game
const (
	Setup GameStage = iota
	Spawn
	SelectActions
	Simulate
)

type Game struct {
	Board Board
	Players []Player
	// Cards
	// Settings
	Stage GameStage
}

func (g *Game) Run() {
	for {
		switch g.Stage {
		case Setup:
			g.setupPhase()
		case Spawn:
			g.spawnPhase()
		case SelectActions:
			g.selectActionsPhase()
		case Simulate:
			g.simulatePhase()
		}
	}
}

func (g *Game) setupPhase() {
	/*
	State:
		Player Ready State
	Actions:
	*/
	// collect players

	// Exit when All players are ready
}

func (g *Game) spawnPhase() {
}

func (g *Game) selectActionsPhase() {
}

func (g *Game) simulatePhase() {
}

type Player struct {
}
