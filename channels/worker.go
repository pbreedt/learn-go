package channels

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	var wg sync.WaitGroup
// 	workChan := make(chan string, 10)
// 	doneChan := make(chan bool, 3)
// 	resultChan := make(chan string)

// 	wg.Add(4)

// 	go generateWork(20, "A", &wg, workChan, doneChan)
// 	go processWork(&wg, workChan, resultChan)
// 	go generateWork(5, "B", &wg, workChan, doneChan)
// 	go processResults(10, &wg, resultChan)
// 	go generateWork(5, "C", &wg, workChan, doneChan)

// 	go closeWhenWorkDone(3, &wg, workChan, doneChan)

// 	wg.Wait()
// 	// close(workChan)
// }

func generateWork(howMuch int, prefix string, wg *sync.WaitGroup, workChan chan<- string, doneChan chan<- bool) {
	defer wg.Done()
	// defer close(workChan)

	for i := 0; i < howMuch; i++ {
		fmt.Printf("new work from %s %d\n", prefix, i)
		workChan <- fmt.Sprintf("%s %d", prefix, i)
	}

	doneChan <- true
}

func processWork(wg *sync.WaitGroup, wc <-chan string, rc chan<- string) {
	defer wg.Done()
	result := ""
	for str := range wc {
		fmt.Println("Processing", str)
		result += "| Got " + str
	}
	rc <- result
}

func processResults(timeout int, wg *sync.WaitGroup, rc <-chan string) {
	defer wg.Done()
	select {
	case res := <-rc:
		fmt.Println("RESULT: ", res)
	case <-time.After(time.Second * time.Duration(timeout)):
		fmt.Println("RESULT: timeout")
	}
}

func closeWhenWorkDone(workProducers int, wg *sync.WaitGroup, workChan chan<- string, doneChan <-chan bool) {
	defer wg.Done()
	producersCompleted := 0

	for {
		<-doneChan
		producersCompleted++
		fmt.Println("Work Producers completed: ", producersCompleted)
		if producersCompleted >= workProducers {
			fmt.Println("All Work Producers completed")
			close(workChan)
			return
		}
	}

}
