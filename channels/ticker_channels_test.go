package channels

import (
	"testing"
	"time"
)

func TestTickerChannel(t *testing.T) {
	tc := TickerChannel{time.Second, 2 * time.Second}

	tc.Run(10 * time.Second)

	t.Log("done")
}
