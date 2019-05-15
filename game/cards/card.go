package cards

import (
	"bytes"
	"encoding/json"

	"github.com/tjbearse/robo/game/coords"
)

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

func (d *Deck) Discard(c Card) {
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
		case Forward:
			for i := 0; i < card.Reps; i++ {
				c = c.ApplyDirectionaly(coords.Offset{0, 1})
				steps = append(steps, c)
			}
		case Backward:
			for i := 0; i < card.Reps; i++ {
				c = c.ApplyDirectionaly(coords.Offset{0, -1})
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

type Command int
const (
	Forward Command = iota
	Backward
	TurnLeft
	TurnRight
	UTurn
)
var commandStr = map[Command]string{
}

// credit: https://gist.github.com/lummie/7f5c237a17853c031a57277371528e87
var toString = map[Command]string{
	Forward: "Forward",
	Backward: "Backward",
	TurnLeft: "TurnLeft",
	TurnRight: "TurnRight",
	UTurn: "UTurn",
}

var toID = map[string]Command{
	"Forward": Forward,
	"Backward": Backward,
	"TurnLeft": TurnLeft,
	"TurnRight": TurnRight,
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
