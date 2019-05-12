package game

// FIXME this deck is very simplified
// there are some circumstance
// when a card could be frozen for a long time on a
// board that aren't accounted for
type Deck struct {
	cards []Card
	pos int
}

func NewDeck(c []Card) *Deck {
	return &Deck{c, 0}
}

func (d *Deck) Deal(n int) ([]Card) {
	cards := make([]Card, 0, n)
	for i:=0; i < n; i++ {
		cards = append(cards, d.cards[d.pos])
		d.pos = (d.pos + 1) % len(d.cards)
	}
	return cards
}

func (d *Deck) Shuffle() {
	d.pos = 0
}

type Command int
const (
	Forward Command = iota
	Backward
	TurnLeft
	TurnRight
	UTurn
)
/* Ideas for expansion
N,E,S,W commands
Strafe commands
*/

type Card struct {
	Priority int
	Command Command
	Reps int
}

// CardCompare is for sorting
func CardCompare(a,b Card) bool {
	return a.Priority < b.Priority
}

func (card Card) Apply(c Configuration) []Configuration {
	steps := make([]Configuration, 0)
	switch card.Command {
		case Forward:
			for i := 0; i < card.Reps; i++ {
				c = c.ApplyDirectionaly(Offset{0, 1})
				steps = append(steps, c)
			}
		case Backward:
			for i := 0; i < card.Reps; i++ {
				c = c.ApplyDirectionaly(Offset{0, -1})
				steps = append(steps, c)
			}
		case TurnLeft:
			c.Heading = c.Heading.RotateRight(-1)
			steps = append(steps, c)
		case TurnRight:
			c.Heading = c.Heading.RotateRight(1)
			steps = append(steps, c)
		case UTurn:
			c.Heading = c.Heading.RotateRight(2)
			steps = append(steps, c)
	}
	return steps
}
