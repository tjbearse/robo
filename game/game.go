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
	// Players
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
}

func (g *Game) spawnPhase() {
}

func (g *Game) selectActionsPhase() {
}

func (g *Game) simulatePhase() {
}

type Player struct {
}
