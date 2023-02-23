package errors

import "fmt"

func basics() {
	err := simpleError()
	if err != nil {
		fmt.Printf("Expected ERR: %v\n", err)
	}

	err = doError(false)
	if err != nil {
		fmt.Printf("Expected ERR: %v\n", err)
	}

	err = doError(true)
	if err != nil {
		fmt.Printf("Expected ERR: %v\n", err)
	}
}

func simpleError() error {
	return fmt.Errorf("simple error")
}

// Anything that implements Error() string is an 'error'
type ErrCode struct {
	code  int
	name  string
	extra interface{}
}

func (e ErrCode) Error() string {
	str := fmt.Sprintf("%04d:%s", e.code, e.name)
	if e.extra != nil {
		return fmt.Sprintf("%s (%s)", str, e.extra)
	}
	return str
}
func (e ErrCode) WithExtra(ex interface{}) error {
	n := e
	n.extra = ex
	return n
}

var (
	Err10 = ErrCode{10, "error 10 occurred", nil}
	Err20 = ErrCode{20, "error 20 occurred", nil}
)

func doError(err bool) error {
	if err {
		return Err10.WithExtra("this is horrible!")
	}
	return nil
}
