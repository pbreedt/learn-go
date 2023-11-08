package channels

import (
	"fmt"
	"sync"
	"time"
)

func basics() {
	runBlocking()
	runWaitGroup()
	runRange()
	runSelect()
	runSelectTime()
}

func multiply(a int, b int, c chan int) {
	time.Sleep(time.Microsecond * time.Duration(a))
	c <- a * b // send result to c
}

func runBlocking() {
	c := make(chan int)
	go multiply(2, 10, c)
	go multiply(3, 5, c)
	x, y := <-c, <-c // receive from c - blocking until both returns

	fmt.Println(x, y)
}

func runWaitGroup() {
	var wg sync.WaitGroup

	c := make(chan int)
	go multiply(2, 10, c)
	go multiply(3, 5, c)

	wg.Add(1)
	var x, y int
	go func() { // start in goroutine
		x, y = <-c, <-c // receive from c - blocking until both returns
		wg.Done()
	}()

	wg.Wait() // wait until goroutine completes with WaitGroup.Done()
	fmt.Println(x, y)
}

func runRange() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		fmt.Printf("fib:%d\n", i)
	}
}

func fibonacci(a int, c chan int) {
	x, y := 0, 1

	for i := 0; i < a; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func runSelect() {
	c := make(chan int)

	for {
		select {
		case x := <-c: // read from c - nothing there, so block
			fmt.Println("net gets x", x)
		case <-time.After(time.Second):
			fmt.Println("timeout")
			return
		}
	}
}

func runSelectTime() {
	tick1 := time.Tick(1500 * time.Millisecond)
	tick := time.NewTicker(1000 * time.Millisecond).C
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-tick1:
			fmt.Println("tick1.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
