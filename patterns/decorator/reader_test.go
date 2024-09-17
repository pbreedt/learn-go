package decorator

import (
	"testing"
)

func TestReader(t *testing.T) {
	fr := &FileReader{}
	cr := &CompressReader{fr}
	er := &EncryptReader{cr}

	data := er.Read()
	expect := "Read file...Compress data...Encrypting data..."
	if data != expect {
		t.Fatalf("Fail: expected:%s, got:%s", expect, data)
	}
}
