package common

// LocalPlayer represents all information known about a player. In a
// server, all players will be of this type. In a client, only the
// local player should use this type.
type LocalPlayer struct {
	ID        int // unique ID assigned by server
	Name      string
	Resources []ResourceType
	// TODO: random cards
	//Cards []
}

// IsLocal returns true for LocalPlayer
func (player *LocalPlayer) IsLocal() bool {
	return true
}
