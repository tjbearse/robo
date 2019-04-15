package game

type GamePhase interface {}

type Game struct {
	board Board
	players map[*Player]bool
	Deck Deck
	// Settings
	phase GamePhase
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

//FIXME, a true type for this
type Message string
type PlayerMessage struct {
	Player *Player
	Message string
}

type Player struct {
	Name string
	Robot Robot
	Spawn *Configuration
}

type Robot struct {
	Name string
	Damage int
	Lives int
	Board []*Card
	Configuration *Configuration
	// TODO stuck cards
}
