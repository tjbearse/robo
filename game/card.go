package game

// FIXME this deck is very simplified
// there are some circumstance
// when a card could be frozen for a long time on a
// board that aren't accounted for
type Deck struct {
	cards []Card
	pos int
}

func NewDeck() Deck {
	return Deck{
		[]Card{
			{ 1, Forward, 1},
			{ 2, Forward, 2},
			{ 3, TurnLeft, 1},
			{ 4, TurnRight, 1},
			{ 5, Backward, 1},
		},
		0,
	}
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
type Card struct {
	Priority int
	Command Command
	Reps int
}
