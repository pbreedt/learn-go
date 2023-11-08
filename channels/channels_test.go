package channels

import (
	"testing"
	"time"
)

func TestBasics(t *testing.T) {
	basics()
}

func TestBlockingWrite(t *testing.T) {
	c := make(chan string, 5)
	go StartReading(c, 1*time.Second)
	go StartWriting(c, 100*time.Millisecond, 15*time.Second)

	time.Sleep(time.Second * 20)
	t.Log("done")
}
