package common

import (
	"math/rand"
	"time"
)

// NewTimeRandom creates a new random number generator from the
// current time
func NewTimeRandom() *rand.Rand {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r
}
