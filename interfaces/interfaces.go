package interfaces

import (
	"fmt"
)

type Noiser interface {
	MakeNoise()
}

type Dog struct{}

func (d *Dog) MakeNoise() {
	fmt.Println("Woof")
}

type Cat struct{}

func (d Cat) MakeNoise() {
	fmt.Println("Meow")
}

func MakeNoiser(typ int) Noiser {
	// Cat can be returned as struct or pointer - Noiser interface implemented by Cat struct/value receiver (not pointer)
	if typ == 1 {
		return Cat{}
	} else if typ == 2 {
		// Cat pointer can be dereferenced to get to Cat value, thus also valid
		return &Cat{}
	} else {
		// must return pointer - Noiser interface implemented by Dog pointer receiver
		return &Dog{}
	}
}

func basics() {

	MakeNoiser(1).MakeNoise()
	MakeNoiser(2).MakeNoise()
	MakeNoiser(3).MakeNoise()

}
