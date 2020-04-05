package common

type ResourceType int

const (
	RESOURCE_BRICKS ResourceType = iota
	RESOURCE_DESERT
	RESOURCE_SHEEP
	RESOURCE_STONE
	RESOURCE_WHEAT
	RESOURCE_WOOD
)

type StructureType int

const (
	STRUCTURE_SETTLEMENT StructureType = iota
	STRUCTURE_CITY
)

// Board stores all information about the board and provides
// high-level access to the board via exposed functions
type Board struct {
	Tiles      [][]Tile
	Structures [][]Structure
	Roads      []Road
}

// Tile represents a single tile on the board
type Tile struct {
	Resource ResourceType
	Number   int
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
