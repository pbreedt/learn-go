package files

import (
	"fmt"
	"io"
)

func Read(r io.Reader) ([]byte, int, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, 0, fmt.Errorf("read error: %v", err)
	}
	return buf, len(buf), nil
}

func Write(w io.Writer, b []byte) (int, error) {
	n, err := w.Write(b)
	if err != nil {
		return 0, fmt.Errorf("write error: %v", err)
	}
	return n, nil
}
