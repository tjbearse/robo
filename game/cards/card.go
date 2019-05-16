package cards

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/tjbearse/robo/game/coords"
)

type Deck struct {
	cards []Card
	// discard tracks both cards in play, i.e. out of deck,
	// and those in the discard pile
	// true denotes "out", false denotes the discard pile
	discard map[Card]bool
}

func NewDeck(c []Card) (*Deck, error) {
	// priorities could be overlapped even if cards aren't
	// so we're gonna check for that
	priorities := map[int]bool{}
	discard := map[Card]bool{}
	for _, card := range(c) {
		if priorities[card.Priority] {
			return nil, errors.New("Deck contains duplicate priorities")
		}
		priorities[card.Priority] = true
		discard[card] = false // discard pile
	}
	// putting all cards in the discard ensures that the deck gets shuffled
	return &Deck{[]Card{}, discard}, nil
}

func (d *Deck) Deal(n int) ([]Card, error) {
	if len(d.cards) < n {
		d.Shuffle()
	}
	if len(d.cards) < n {
		return []Card{}, fmt.Errorf("can't deal desired cards: have %d, want %d", len(d.cards), n)
	}
	cards := d.cards[:n]
	d.cards = d.cards[n:]

	// add to discard
	for _, card := range(cards) {
		d.discard[card] = true
	}
	return cards, nil
}

func (d *Deck) Shuffle() {
	// should always append cards to existing cards
	fresh := make([]Card, 0, len(d.discard))
	for c, inPlay := range(d.discard) {
		if !inPlay {
			fresh = append(fresh, c)
		}
	}
	rand.Shuffle(len(fresh), func(i, j int) {
		fresh[i], fresh[j] = fresh[j], fresh[i]
	})
	d.cards = append(d.cards, fresh...)
}

func (d *Deck) Discard(c Card) (error) {
	if !d.discard[c] {
		return fmt.Errorf("card not listed as in play, %v", c)
	}
	d.discard[c] = false
	return nil
}


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

func (card Card) Apply(c coords.Configuration) []coords.Configuration {
	steps := make([]coords.Configuration, 0)
	switch card.Command {
		case Move:
			for i := 0; i < card.Reps; i++ {
				c = c.ApplyDirectionaly(coords.Offset{0, 1})
				steps = append(steps, c)
			}
		case BackUp:
			for i := 0; i < card.Reps; i++ {
				c = c.ApplyDirectionaly(coords.Offset{0, -1})
				steps = append(steps, c)
			}
		case RotateLeft:
			c.Heading = c.Heading.RotateRight(-1)
			steps = append(steps, c)
		case RotateRight:
			c.Heading = c.Heading.RotateRight(1)
			steps = append(steps, c)
		case UTurn:
			c.Heading = c.Heading.RotateRight(2)
			steps = append(steps, c)
	}
	return steps
}

type Command int
const (
	Move Command = iota
	BackUp
	RotateLeft
	RotateRight
	UTurn
)
var commandStr = map[Command]string{
}

// credit: https://gist.github.com/lummie/7f5c237a17853c031a57277371528e87
var toString = map[Command]string{
	Move: "Move",
	BackUp: "BackUp",
	RotateLeft: "RotateLeft",
	RotateRight: "RotateRight",
	UTurn: "UTurn",
}

var toID = map[string]Command{
	"Move": Move,
	"BackUp": BackUp,
	"RotateLeft": RotateLeft,
	"RotateRight": RotateRight,
	"UTurn": UTurn,
}
// MarshalJSON marshals the enum as a quoted json string
func (c Command) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[c])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (c *Command) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*c = toID[j]
	return nil
}
