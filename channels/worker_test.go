package channels

import (
	"sync"
	"testing"
)

func TestWorkerChannel(t *testing.T) {
	var wg sync.WaitGroup
	workChan := make(chan string, 10)
	doneChan := make(chan bool, 3)
	resultChan := make(chan string)

	wg.Add(4)

	go generateWork(20, "A", &wg, workChan, doneChan)
	go processWork(&wg, workChan, resultChan)
	go generateWork(5, "B", &wg, workChan, doneChan)
	go processResults(10, &wg, resultChan)
	go generateWork(5, "C", &wg, workChan, doneChan)

	go closeWhenWorkDone(3, &wg, workChan, doneChan)

	wg.Wait()
	// close(workChan)
}
