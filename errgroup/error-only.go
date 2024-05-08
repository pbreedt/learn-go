package errgroup

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// errgroup.Group is related to sync.WaitGroup but adds handling of tasks returning errors.

func runErrGroup() {
	g := new(errgroup.Group)
	var payloads = []string{
		"this work",
		"that work",
		"some other work",
	}
	for _, work := range payloads {
		// Launch a goroutine to do work in parallel
		w := work // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			err := doWork(w)
			if err != nil {
				fmt.Printf("Error doing work '%s': %v\n", w, err)
			}
			return err
		})
	}
	// Wait for all work to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully completed all work.")
	}
}

func doWork(work string) error {
	fmt.Println("Doing work:", work)
	time.Sleep(200 * time.Millisecond)
	return nil
}
