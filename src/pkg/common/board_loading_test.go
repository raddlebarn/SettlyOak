package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSimpleTextTestMap(t *testing.T) {
	const (
		F = false
		T = true
	)
	var (
		expectedXSize = 5
		expectedYSize = 5
		expectedTiles = [][]bool{
			{F, F, T, T, T},
			{F, T, T, T, T},
			{T, T, T, T, T},
			{T, T, T, T, F},
			{T, T, T, F, F},
		}
	)

	board, err := LoadTxtMap("testmap1.txt")

	assert.Equal(t, nil, err, "An error should not occur")

	// check dimensions
	for y := 0; y < len(board.Tiles); y++ {
		assert.Equal(t, expectedXSize, len(board.Tiles[0]),
			fmt.Sprintf("Board should be %d tiles wide",
				expectedXSize),
		)
	}
	assert.Equal(t, expectedYSize, len(board.Tiles),
		fmt.Sprintf("Board should be %d tiles tall",
			expectedXSize),
	)

	// verify all the tiles
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			fmt.Printf("%d", board.Tiles[y][x].Resource)
			assert.Equal(t, expectedTiles[y][x], board.Tiles[y][x].Resource != RESOURCE_BLANK, "Tile did not match expectation")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
