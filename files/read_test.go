package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	b, n, e := ReadFile("test-file.txt")
	// fmt.Printf("read %d bytes: %s (error=%v)", n, b, e)
	assert.Equal(t, nil, e)
	assert.Equal(t, []byte("abc\ndef\ngh"), b)
	assert.Equal(t, 10, n)
}
