package errors

import (
	"testing"
)

func TestCustomError(t *testing.T) {
	e := ThrowErr(ErrDatabase)
	if e != ErrDatabase {
		t.Fatalf("incorrect error. expect:%v, got:%v", ErrDatabase, e)
	}
	e = ThrowErr(ErrInvalidInput)
	if e != ErrInvalidInput {
		t.Fatalf("incorrect error. expect:%v, got:%v", ErrInvalidInput, e)
	}
	e = ThrowErr(ErrUserNotFound)
	if e != ErrUserNotFound {
		t.Fatalf("incorrect error. expect:%v, got:%v", ErrUserNotFound, e)
	}
}

func ThrowErr(ce CustomError) error {
	return ce
}
