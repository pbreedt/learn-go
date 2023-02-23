package testing

import (
	"fmt"
	"testing"
)

// computing test coverage
// and displaying the green/red source
// go test -coverprofile=cover.out
// go tool cover -html=cover.out

// benchmarking time and memory allocations
// go test -bench=.
// go test -bench=. -benchmem
func Lengthy() {
	for i := 0; i < 10000000; i++ {
	}
}

func BenchmarkLengthy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lengthy()
	}
}

// profiling CPU, memory and blocking
// go test -bench=. -cpuprofile=cpu.out
// go test -bench=. -memprofile=mem.out
// go test -bench=. -blockprofile=block.out
// go tool pprof -text -nodecount=10 ./cpu.out

// providing examples
// included in the documentation
// output checked when tests are run
func Division(x, y int) int {
	return x / y
}

func ExampleDivision() {
	fmt.Println(Division(4, 2))
	fmt.Println(Division(10, 2))
	// Output:
	// 2
	// 5
}
