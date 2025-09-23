package files

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	r := strings.NewReader("abcde")
	b, n, e := Read(r)
	assert.Equal(t, nil, e)
	assert.Equal(t, []byte("abcde"), b)
	assert.Equal(t, 5, n)
}

func TestWrite(t *testing.T) {
	var sb strings.Builder
	n, e := Write(&sb, []byte("hello"))
	assert.Equal(t, nil, e)
	assert.Equal(t, 5, n)
	assert.Equal(t, "hello", sb.String())
}
