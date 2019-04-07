package helperfunc

import (
	"crypto/rand"
	"fmt"
)

// helper functions

func Tokengenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
