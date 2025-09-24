package race

import (
	"fmt"
	"sync"
)

func Race() {
	i := 0
	go func() {
		i++ // write
	}()
	i++
	fmt.Println(i) // concurrent read
}

func NoRace() {
	c := make(chan int)
	go func() {
		locali := 0
		locali++ // write
		c <- locali
	}()

	i := <-c
	i++
	fmt.Println(i) // no concurrent read
}

func NoRaceMutex() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1) // write
		close(wait)
	}()
	n.Add(1) // write, no race condition
	<-wait
	fmt.Println(n.Value())
}

type AtomicInt struct {
	mu sync.Mutex
	n  int
}

// always use pointer receiver for mutex (mutex should never be copied)
func (a *AtomicInt) Add(n int) {
	a.mu.Lock()
	a.n += n
	a.mu.Unlock()
}

func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}
