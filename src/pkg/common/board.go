package common

import (
	"errors"
	"math/rand"
)

type ResourceType int

const (
	RESOURCE_BLANK  ResourceType = iota // 0
	RESOURCE_DESERT                     // 1
	RESOURCE_BRICKS                     // 2
	RESOURCE_SHEEP                      // 3
	RESOURCE_STONE                      // 4
	RESOURCE_WHEAT                      // 5
	RESOURCE_WOOD                       // 6
	resourceEnd                         // used for random
)

// ErrInvalidResource is the error returned if the resource number is
// less than 0 or exceeds resourceEnd
var ErrInvalidResource = errors.New("invalid resource number")

type StructureType int

const (
	STRUCTURE_SETTLEMENT StructureType = iota
	STRUCTURE_CITY
)

// Board stores all information about the board and provides
// high-level access to the board via exposed functions
type Board struct {
	Tiles      [][]*Tile      // Tiles[y][x]
	Structures [][]*Structure // Structures[y][x]
	Roads      []*Road

	rand *rand.Rand
}

// NewBoard creates a new, empty board with no tiles
func NewBoard() *Board {
	return &Board{
		Tiles:      [][]*Tile{},
		Structures: [][]*Structure{},
		Roads:      []*Road{},

		rand: NewTimeRandom(),
	}
}

// Tile represents a single tile on the board
type Tile struct {
	Resource ResourceType
	Number   int
}

// NewTile creates a new tile of a specific type. Set number to 0 to
// assign it later.
func NewTile(resource ResourceType, number int) (*Tile, error) {
	if resource < 0 || resource >= resourceEnd {
		return nil, ErrInvalidResource
	}

	return &Tile{
		Resource: resource,
		Number:   number,
	}, nil
}

// Structure represents a structure built on a tile vertex
type Structure struct {
	Type  StructureType
	Owner *Player
}

// Road represents a road built on a tile edge
type Road struct {
	Start int
	End   int
	Owner *Player
}
