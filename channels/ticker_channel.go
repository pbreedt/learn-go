package channels

import (
	"fmt"
	"time"
)

type TickerChannel struct {
	Interval  time.Duration
	WorkDelay time.Duration
}

func (tc TickerChannel) Run(d time.Duration) {
	ticker := time.NewTicker(tc.Interval)
	counter := 0

	time.AfterFunc(d, func() { ticker.Stop() }) // can't close ticker.C here, chan is read-only

	for {
		select { // cannot use range: ticker.C not getting closed, hence the timeout
		case t := <-ticker.C:
			counter++
			go tc.DoSomething(t, counter)
			go func(tt time.Time) {
				fmt.Printf("%s: counter=%d: Ticker ---done-- %s\n", time.Now().Format(time.RFC3339Nano), counter, tt.Format(time.RFC3339Nano))
			}(t)
		case <-time.After(tc.Interval + time.Second):
			return
		}
	}
}

func (tc TickerChannel) DoSomething(t time.Time, c int) {
	fmt.Printf("%s: counter=%d: Ticker  received %s\n", time.Now().Format(time.RFC3339Nano), c, t.Format(time.RFC3339Nano))
	time.Sleep(tc.WorkDelay)
	fmt.Printf("%s: counter=%d: Ticker processed %s\n", time.Now().Format(time.RFC3339Nano), c, t.Format(time.RFC3339Nano))
}
