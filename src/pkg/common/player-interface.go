package common

// Player interface represents any kind of player, and will expose
// different degrees of information dependent on whether they are
// local or not
type Player interface {
	// IsLocal returns true if this is a LocalPlayer
	IsLocal() bool
}
