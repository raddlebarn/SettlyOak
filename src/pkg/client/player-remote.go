package common

// RemotePlayer represents a player with unknown information,
// e.g. unkown inventory, due to not being the current player. Should
// not be used by the server, as the server knows all.
type RemotePlayer struct {
	ID   int // unique ID assigned by server
	Name string
}

// IsLocal returns false for RemotePlayer
func (player *RemotePlayer) IsLocal() bool {
	return false
}
