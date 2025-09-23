package files

import (
	"fmt"
	"os"
)

func WriteNewFile(filename string, data []byte) (int, error) {
	f, e := os.Create(filename) // truncates if file exists
	if e != nil {
		fmt.Printf("error creating file: %v", e)
		return 0, e
	}

	n, e := f.Write(data)
	return n, e
}

func WriteExistingFile(filename string, data []byte) (int, error) {
	f, e := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644) // does not truncate (os.O_APPEND?)
	if e != nil {
		fmt.Printf("error opening file: %v", e)
		return 0, e
	}

	n, e := f.Write(data)
	return n, e
}
