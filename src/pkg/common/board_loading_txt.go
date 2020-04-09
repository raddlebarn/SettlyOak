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

// LoadPreset loads a preset from 2D slices.  "numbers" is
// optional -- use an empty slice {{}} if you don't want to specify
func (board *Board) LoadPreset(tiles [][]byte, numbers [][]int) error {
	board.rand = NewTimeRandom()

	// reset tiles
	board.Tiles = [][]*Tile{{}}

	// iterate
	var y int
	for y = 0; y < len(tiles); y++ {
		for x := 0; x < len(tiles[y]); x++ {
			char := tiles[y][x]
			var number int
			if y < len(numbers) && x < len(numbers[y]) {
				number = numbers[y][x]
			}
			switch char {
			case '\n':
				break
			case '.':
				newTile, _ := NewTile(0, number)
				board.Tiles[y] = append(board.Tiles[y],
					newTile)

			// TODO: for these two cases, randomly
			// generate the numbers with good weightings
			case 'x':
				newTile := board.newRandomTile(number, false)
				board.Tiles[y] = append(board.Tiles[y],
					newTile)

			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				resource := ResourceType(char - 48)
				newTile, err := NewTile(resource, number)
				if err != nil {
					return err
				}
				board.Tiles[y] = append(board.Tiles[y],
					newTile)

			default:
				return ErrInvalidChar
			}

			// finished the row; add a new one
			board.Tiles = append(board.Tiles, []*Tile{})
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
func LoadTxtMap(fp string) (*Board, error) {
	// read entire file into a buffer as this is just a map preset
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	// convert data to [][]byte of tile characters
	tiles := make([][]byte, 0)
	currentRow := make([]byte, 0)
	for _, char := range data {
		switch char {
		case '\n':
			tiles = append(tiles, currentRow)
			currentRow = make([]byte, 0)
		default:
			currentRow = append(currentRow, char)
		}
	}
	if len(currentRow) > 0 {
		tiles = append(tiles, currentRow)
	}

	board := NewBoard()
	err = board.LoadPreset(tiles, make([][]int, 0))
	if err != nil {
		return nil, err
	}
	return board, nil
}

// newRandomTile produces a random Tile from all the available
// resources. The desert bool controls whether deserts are
// allowed. Tile.Number is not set. Panics if the generated tile is
// invalid.
// TODO: abstract this to a specific generator
func (board *Board) newRandomTile(number int, desert bool) *Tile {
	lower := int(RESOURCE_BRICKS)
	if desert {
		lower = int(RESOURCE_DESERT)
	}
	upper := int(resourceEnd)

	n := upper - lower
	r := board.rand.Intn(n) + lower

	// TODO: if number == 0, generate a random number
	tile, err := NewTile(ResourceType(r), number)
	if err != nil {
		log.Panicf("error generating new random tile: %s", err)
	}
	return tile
}
