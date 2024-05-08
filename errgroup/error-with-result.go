package errgroup

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// errgroup.Group is related to sync.WaitGroup but adds handling of tasks returning errors.
// sample straight from https://pkg.go.dev/golang.org/x/sync/errgroup

type Result string

func runErrGroupWithResult() {
	var g errgroup.Group

	var payloads = []string{
		"this work",
		"that work",
		"some other work",
	}
	results := make([]Result, len(payloads))
	for idx, work := range payloads {
		i, w := idx, work // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			result, err := doWorkWithResult(w)
			if err == nil {
				results[i] = result
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully completed all work.")
	}

	for _, result := range results {
		fmt.Println(result)
	}

}

func doWorkWithResult(work string) (Result, error) {
	fmt.Println("Doing work:", work)
	time.Sleep(200 * time.Millisecond)
	return Result("Work done: " + work), nil
}
