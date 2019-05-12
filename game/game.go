package game

type GamePhase interface {}

type Game struct {
	Board Board
	players map[*Player]bool
	Deck Deck
	// Settings
	phase GamePhase
}

func NewGame(board Board, deck Deck, initialPhase GamePhase) Game {
	return Game{
		board,
		map[*Player]bool{},
		deck,
		initialPhase,
	}
}

func (g Game) GetPlayers() (map[*Player]bool) {
	return g.players
}
func (g *Game) UpdatePlayers(players map[*Player]bool) {
	g.players = players
}
func (g *Game) GetPhase() GamePhase {
	return g.phase
}
func (g *Game) ChangePhase(ph GamePhase) {
	g.phase = ph
}
func (g *Game) GetNextSpawn() (Configuration, error){
	return g.Board.getNextSpawn()
}

func (g *Game) CheckForRobot(c Coord) *Robot {
	for p, _ := range(g.players) {
		if p.Robot.Configuration != nil && p.Robot.Configuration.Location == c {
			return &p.Robot
		}
	}
	return nil
}

type Player struct {
	Name string
	Robot Robot
	Spawn Spawn
	FlagNum int // i.e. the flag currently targetting
}

type Robot struct {
	Name string
	Damage int
	Lives int
	Board []*Card
	Configuration *Configuration
}

type SpawnState int
const (
	Unset SpawnState = iota
	Rotatable
)
type Spawn struct {
	State SpawnState
	Coord Coord
}
