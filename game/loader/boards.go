package loader

import (
	"io"
	"bufio"
	"strings"
	"strconv"
	"errors"
	"fmt"

	. "github.com/tjbearse/robo/game"
	. "github.com/tjbearse/robo/game/coords"
	"github.com/tjbearse/robo/game/coords"
)

// TODO build a serializer here and copy some of the known board
// would also need a way to select in create game


var DimError = errors.New("couldn't read dimensions")

func GetBoard(reader *bufio.Reader) (*Board, error) {
	width, height, err := processDim(reader)
	if err != nil {
		return nil, err
	}
	pb := newPlainBoard(width, height)
	for x := 0; x < width; x++ {
		swall, _, err := getWallBurn(reader)
		if err != nil {
			err = wrapError(err, fmt.Sprintf("failed wall @ %d,%d", x, 0))
			return nil, err
		}
		pb.Nwalls[x][0] = swall
	}
	err = ensureNewLine(reader)
	if err != nil {
		return nil, wrapError(err, "parsing wall line")
	}
	spawns := map[int]coords.Configuration{}
	flags := map[int]coords.Coord{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tile, num, swall, wwall, err := getTile(reader)
			if err != nil {
				err = wrapError(err, fmt.Sprintf("failed tile %d,%d", x, y))
				return nil, err
			}
			pb.Tiles[x][y] = tile
			pb.Wwalls[x][y] = wwall
			pb.Nwalls[x][y+1] = swall
			if tile.Type == Spawn {
				spawns[num] = coords.Configuration{coords.Coord{x,y}, North}
			} else if tile.Type == Flag {
				flags[num] = coords.Coord{x,y}
			}
		}
		_, wwall, err := getWallOnly(reader)
		if err != nil {
			err = wrapError(err, fmt.Sprintf("failed wall @ %d,%d", width, y))
			return nil, err
		}
		pb.Wwalls[width][y] = wwall
		err = ensureNewLine(reader)
		if err != nil {
			err = wrapError(err, fmt.Sprintf("failed line %d", y+1))
			return nil, err
		}
	}
	pb.FlagOrder = make([]coords.Coord, len(flags))
	for i:=0; i<len(flags); i++ {
		if _, ok := flags[i+1]; !ok {
			return nil, errors.New("flag numbers differed from expected")
		}
		pb.FlagOrder[i] = flags[i+1]
	}
	spawnSlice := make([]coords.Configuration, len(spawns))
	for i:=0; i<len(spawns); i++ {
		if _, ok := spawns[i+1]; !ok {
			return nil, fmt.Errorf("spawn numbers differed from expected: i=%d %v", i, spawns)
		}
		spawnSlice[i] = spawns[i+1]
	}

	return NewBoard(pb, spawnSlice)
}

func processDim(reader *bufio.Reader) (width int, height int, err error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		err = wrapError(err, "reading dimensions")
		return
	}
	dim := strings.Fields(str)
	if len(dim) != 2 {
		return 0, 0, DimError
	}
	height, err = strconv.Atoi(dim[0])
	if err != nil {
		err = wrapError(err, "reading height")
		return
	}
	width, err = strconv.Atoi(dim[1])
	if err != nil {
		err = wrapError(err, "reading width")
		return
	}
	return
}

func getWallOnly(reader *bufio.Reader) (sw, ww bool, err error) {
	buf := make([]byte, 1)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		err = wrapError(err, "couldn't read a byte for wall")
		return
	}
	return parseWall(buf[0])
}

func getWallBurn(reader *bufio.Reader) (sw, ww bool, err error) {
	buf := make([]byte, 3)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		err = wrapError(err, "couldn't read 3 bytes for wall")
		return
	}
	return parseWall(buf[0])
}

func getTile(reader *bufio.Reader) (t Tile, num int, sw bool, ww bool, err error) {
	buf := make([]byte, 3)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		err = wrapError(err, "couldn't read 3 bytes for tile")
		return
	}
	wB := buf[0]
	tB := buf[1]
	tB2 := buf[2]

	sw, ww, err = parseWall(wB)
	if err != nil {
		return
	}

	t, num, err = parseTile(tB, tB2)
	return
}

func parseWall(b byte) (sw, ww bool, err error) {
	switch b {
	case '|':
		ww = true
	case '_':
		sw = true
	case 'L':
		sw = true
		ww = true
	case ' ':
	default:
		err = fmt.Errorf("unexpected wall marker, %q", b)
	}
	return
}


func parseTile(a, b byte) (t Tile, num int, err error) {
	switch a {
	case 'R':
		t.Type = Repair
	case 'X':
		t.Type = Pit
	case 'U':
		t.Type = Upgrade
	case 'g':
		t.Type = Gear
	case 'G':
		t.Type = Gear
		t.Dir = South
	case 'P':
		t.Type = Pusher
		t.Dir = toDir(b)
	case 'L':
		t.Type = Laser
		t.Dir = toDir(b)
	case 'F':
		t.Type = Flag
		num, err = strconv.Atoi(string(b))
	case 'S':
		t.Type = Spawn
		t.Dir = North
		num, err = strconv.Atoi(string(b))
	case '<':
		fallthrough
	case 'v':
		fallthrough
	case '^':
		fallthrough
	case '>':
		t = toConveyor(a, b)
	case '.':
	default:
		err = fmt.Errorf("unexpected tile marker, %q,%q", a, b)
	}
	return
}

func toDir(b byte) (Dir) {
	switch b {
	case 's':
		return South
	case 'w':
		return West
	case 'e':
		return East
	default:
		return North
	}
}

func toConveyor(a, b byte) (t Tile) {
	if a == b {
		t.Type = ExpressConveyor
	} else {
		t.Type = Conveyor
	}
	switch a {
	case '<':
		t.Dir = West
	case '>':
		t.Dir = East
	case '^':
		t.Dir = North
	case 'v':
		t.Dir = South
	}
	return
}

func newPlainBoard(width, height int) (PlainBoard) {
	pb := PlainBoard{
		Tiles: make([][]Tile, width),
		Nwalls: make([][]bool, width+1),
		Wwalls: make([][]bool, width+1),
		FlagOrder: []coords.Coord{},
	}
	x := 0
	for x < width {
		pb.Tiles[x] = make([]Tile, height)
		pb.Nwalls[x] = make([]bool, height+1)
		pb.Wwalls[x] = make([]bool, height+1)
		x++
	}
	pb.Nwalls[x] = make([]bool, height+1)
	pb.Wwalls[x] = make([]bool, height+1)
	return pb
}

func ensureNewLine(reader *bufio.Reader) (error) {
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		return err
	}
	for _, b := range(bytes) {
		if b != '\n' && b != ' ' {
			return fmt.Errorf("found chars while clearing new line, %q", bytes)
		}
	}
	return nil
}

func wrapError(e error, reason string) (error) {
	return fmt.Errorf("%s: %s", reason, e.Error())
}
