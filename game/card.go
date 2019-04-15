package game

// FIXME this deck is very simplified
// there are some circumstance
// when a card could be frozen for a long time on a
// board that aren't accounted for
type Deck struct {
	cards []Card
	pos int
}

func (d *Deck) Deal(n int) ([]Card) {
	cards := make([]Card, n)
	for i:=0; i < n; i++ {
		cards = append(cards, d.cards[d.pos])
		d.pos = (d.pos + 1) % len(cards)
	}
	return cards
}

func (d *Deck) Shuffle() {
	// TODO
}

type Card struct {
	Priority int
	// Command
}
