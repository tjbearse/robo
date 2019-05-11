package game

type GamePhase interface {}

type Game struct {
	board Board
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
	return g.board.getNextSpawn()
}

type Player struct {
	Name string
	Robot *Robot
	Spawn Spawn
	// TODO hand?
}

type SpawnState int
const (
	Unset SpawnState = iota
	Fixed
	Rotatable
)
type Spawn struct {
	State SpawnState
	Config Configuration 
}

type Robot struct {
	Name string
	Damage int
	Lives int
	Board []*Card
	Configuration *Configuration
	// TODO stuck cards
}
