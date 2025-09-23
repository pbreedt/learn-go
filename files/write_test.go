package files

import (
	"testing"
)

func TestWriteNewFile(t *testing.T) {
	filename := "test_new.txt"
	data := []byte("Hello, World!")
	n, err := WriteNewFile(filename, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if n != len(data) {
		t.Fatalf("Expected %d bytes written, got %d", len(data), n)
	}
}

func TestWriteExistingFile(t *testing.T) {
	filename := "test_existing.txt"
	initialData := []byte("Initial Data\n")
	_, err := WriteNewFile(filename, initialData)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	moreData := []byte("More Data\n")
	n, err := WriteExistingFile(filename, moreData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if n != len(moreData) {
		t.Fatalf("Expected %d bytes written, got %d", len(moreData), n)
	}
	n, err = WriteExistingFile(filename, moreData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if n != len(moreData) {
		t.Fatalf("Expected %d bytes written, got %d", len(moreData), n)
	}
}
