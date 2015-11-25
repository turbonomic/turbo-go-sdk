package rand

import (
	"math/rand"
	"sync"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
var numLetters = len(letters)
var rng = struct {
	sync.Mutex
	rand *rand.Rand
}{
	rand: rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
}

// String generates a random alphanumeric string n characters long.  This will
// panic if n is less than zero.
func String(n int) string {
	if n < 0 {
		panic("out-of-bounds value")
	}
	b := make([]rune, n)
	rng.Lock()
	defer rng.Unlock()
	for i := range b {
		b[i] = letters[rng.rand.Intn(numLetters)]
	}
	return string(b)
}

// Seed seeds the rng with the provided seed.
func Seed(seed int64) {
	rng.Lock()
	defer rng.Unlock()

	rng.rand = rand.New(rand.NewSource(seed))
}
