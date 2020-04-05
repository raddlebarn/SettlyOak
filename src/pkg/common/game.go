package common

// Game stores all information about the current game
type Game struct {
	Board Board

	Players   []Player
	WhoseTurn int
}
