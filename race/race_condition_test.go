package race

import (
	"testing"
)

// test for race condition:
// go test . -v -race
func TestMain(m *testing.T) {
	// Race()
	// NoRace()
	NoRaceMutex()
}
