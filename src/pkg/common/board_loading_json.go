package common

import (
	"encoding/json"
	"os"
)

// JSONMap encodes the information for the JSON-format maps. This is
// extensible without issue, as pretty much only the tiles field is
// required
type JSONMap struct {
	Tiles   [][]byte `json:"tiles"`
	Numbers [][]int  `json:"numbers,omitempty"`
}

func LoadJSONMap(fp string) (*Board, error) {
	// open file for the decoder
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	// decode the json
	var m JSONMap
	decoder := json.NewDecoder(file)
	decoder.Decode(&m)

	// TODO: convert JSONMap to a board
	return board, nil
}
