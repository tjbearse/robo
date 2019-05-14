package coords

type Coord struct {
	X int
	Y int
}

type Configuration struct {
	Location Coord
	Heading Dir
}

type Offset Coord

// Directional offsets assume that forward is south (increasing numbers)

func (o Offset) Rotate(d Dir) Offset {
	switch d {
		case North:
		o.X, o.Y = -o.X, -o.Y
		case East:
		o.X, o.Y = o.Y, -o.X
		case South:
		case West:
		o.X, o.Y = -o.Y, o.X
	}
	return o
}

func (c Coord) Apply(o Offset) Coord {
	c.X += o.X
	c.Y += o.Y
	return c
}

func (c Coord) OffsetTo(b Coord) Offset {
	return Offset{
		b.X - c.X,
		b.Y - c.Y,
	}
}

func (c Configuration) ApplyDirectionaly(o Offset) Configuration {
	c.Location = c.Location.Apply(o.Rotate(c.Heading))
	return c
}
