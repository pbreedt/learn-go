package channels

import (
	"fmt"
	"time"
)

// StartReading reads all messages from channel every rInterval
// Not trying to keep channel empty, want to open some space, but then want it to fill up again
// Want writer to block occasionally
func StartReading(c chan string, rInterval time.Duration) {
	t := time.NewTicker(rInterval)
	for {
		<-t.C
		qLen := len(c)
		fmt.Println("Read tick", "Q len", qLen)
		for i := 0; i < qLen; i++ {
			fmt.Println("Read value", <-c)
		}
		fmt.Println("Read Q cleared")
	}
}

// StartWriting initiates a write to channel ever wInterval, keep doing this for duration d
func StartWriting(c chan string, wInterval time.Duration, d time.Duration) {
	fmt.Println("Start writing")

	t := time.NewTicker(wInterval)
	stopTime := time.Now().Add(d)
	for ok := time.Now().Before(stopTime); ok; {
		<-t.C
		fmt.Println("Write tick")
		write(c)
	}
	fmt.Println("Done writing")

	t.Stop()
}

// actual writing to channel
func write(c chan string) {
	s := time.Now().GoString()
	fmt.Println("Writing", s)
	select {
	// write to channel in select, so we can add a timeout when channel is full
	case c <- s:
		fmt.Println("Wrote to Q", s)
	// timeout handles full channel, delays another write attempt that will block until it can write
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timeout writing to Q", s)
		time.AfterFunc(10*time.Second, func() {
			fmt.Println("REWrite to Q, on resubmit", s)
			c <- s
		})
	}
}
