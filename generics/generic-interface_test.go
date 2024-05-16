package generics

import (
	"reflect"
	"testing"
)

func TestGenericStruct(t *testing.T) {
	strGetter := NewGetter[string]()
	strVal := strGetter.Get()

	if reflect.TypeOf(strVal).Kind() != reflect.String {
		t.Fatalf("expected 'string', got:%T", strVal)
	}

	intGetter := NewGetter[int]()
	intVal := intGetter.Get()

	if reflect.TypeOf(intVal).Kind() != reflect.Int {
		t.Fatalf("expected 'int', got:%T", intVal)
	}
}

func TestStringStruct(t *testing.T) {
	strGetter := NewStringGetter()
	strVal := strGetter.Get()

	if reflect.TypeOf(strVal).Kind() != reflect.String {
		t.Fatalf("expected 'string', got:%T", strVal)
	}
}
