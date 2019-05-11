package game

import (
	"encoding/json"
	"bytes"
)

// Dir describes the facing of an object
type Dir int
// Directions, values are in numerical order clockwise
const (
	North Dir = 0
	East Dir = iota
	South
	West
)

func (d Dir) String() string {
	return toString[d]
}

// credit: https://gist.github.com/lummie/7f5c237a17853c031a57277371528e87
var toString = map[Dir]string{
	North: "North",
	East: "East",
	South: "South",
	West: "West",
}

var toID = map[string]Dir{
	"North": North,
	"East": East,
	"South": South,
	"West": West,
}

// MarshalJSON marshals the enum as a quoted json string
func (d Dir) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[d])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (d *Dir) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*d = toID[j]
	return nil
}
