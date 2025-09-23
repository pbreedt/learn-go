package files

import (
	"fmt"
	"os"
)

func ReadFile(filename string) ([]byte, int, error) {
	f, e := os.Open(filename)
	if e != nil {
		fmt.Printf("error opening file: %v", e)
		return nil, 0, e
	}

	b := make([]byte, 10)
	n, e := f.Read(b)
	return b, n, e
}
