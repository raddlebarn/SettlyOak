package common

import (
	"errors"
	"io/ioutil"
	"log"
)

/* # Preset Format Cheat Sheet #

Hexagonal tiles are addressed in an (x,y) system and stored in a
Cartesian grid.
Map presets are plaintext.

Key:
 .   = nothing/sea (synonymous with 0)
 x   = random non-sea, non-desert resource
 0-6 = specific resource (see board.go)

*/

// ErrInvalidChar is the error raised if an invalid character is
// encountered in a map preset
var ErrInvalidChar = errors.New("invalid character")

// LoadPreset loads a preset from a []byte
func (board *Board) LoadPreset(data []byte) error {
	board.rand = NewTimeRandom()

	// reset tiles
	board.Tiles = [][]*Tile{{}}

	// iterate through data
	var y int
	for _, char := range data {
		switch char {
		case '\n':
			y += 1
			board.Tiles = append(board.Tiles, []*Tile{})
		case '.':
			newTile, _ := NewTile(0, 0)
			board.Tiles[y] = append(board.Tiles[y],
				newTile)

		// TODO: for these two cases, randomly generate the
		// numbers with good weightings
		case 'x':
			newTile := board.newRandomTile(false)
			board.Tiles[y] = append(board.Tiles[y],
				newTile)

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			resource := ResourceType(char - 48)
			newTile, err := NewTile(resource, 0)
			if err != nil {
				return err
			}
			board.Tiles[y] = append(board.Tiles[y],
				newTile)

		default:
			return ErrInvalidChar
		}
	}

	// remove any empty bottom row(s)
	y = len(board.Tiles) - 1
	for len(board.Tiles[y]) == 0 {
		board.Tiles = board.Tiles[:y]
		y -= 1
	}

	return nil
}

// LoadTxtMap loads a board from a .txt map file. This does not have
// the ability to specify any map generation parameters.
func (board *Board) LoadTxtMap(fp string) error {
	// read entire file into a buffer as this is just a map preset
	// TODO: support loading entire scenarios
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	return board.LoadPreset(data)
}

// newRandomTile produces a random Tile from all the available
// resources. The desert bool controls whether deserts are
// allowed. Tile.Number is not set. Panics if the generated tile is
// invalid.
func (board *Board) newRandomTile(desert bool) *Tile {
	lower := int(RESOURCE_BRICKS)
	if desert {
		lower = int(RESOURCE_DESERT)
	}
	upper := int(resourceEnd)

	n := upper - lower
	r := board.rand.Intn(n) + lower

	tile, err := NewTile(ResourceType(r), 0)
	if err != nil {
		log.Panicf("error generating new random tile: %s", err)
	}
	return tile
}
