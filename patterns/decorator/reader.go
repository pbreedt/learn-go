package decorator

import "fmt"

type Reader interface {
	Read() string
}

type FileReader struct{}

func (fr *FileReader) Read() string {
	return "Read file..."
}

type CompressReader struct {
	reader Reader
}

func (cr *CompressReader) Read() string {
	file := cr.reader.Read()
	return file + "Compress data..."
}

type EncryptReader struct {
	reader Reader
}

func (er *EncryptReader) Read() string {
	file := er.reader.Read()
	return file + "Encrypting data..."
}

func main() {
	fr := &FileReader{}
	cr := &CompressReader{fr}
	er := &EncryptReader{cr}

	fmt.Println(er.Read())
}
