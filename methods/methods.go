package methods

import "fmt"

type MyStruct struct {
	str string
	i   int
}

// method: pointer receiver can receive both pointer and struct
func (s *MyStruct) SetInt(i int) {
	s.i = i
}

// function: pointer argument MUST be pointer
func SetIntFunc(s *MyStruct, i int) {
	s.i = i
}

// method: struct receiver can also receive both pointer and struct
func (s MyStruct) Print() {
	fmt.Printf("str:%s | int:%d\n", s.str, s.i)
}

// function: struct argument MUST be struct
func Print(s MyStruct) {
	fmt.Printf("str:%s | int:%d\n", s.str, s.i)
}

func basics() {
	s := MyStruct{"one", 1}
	p := &s
	s.SetInt(2)        // methods do auto deref: can use MyStruct or pointer
	s.Print()          // can use MyStruct or pointer for method receiver
	SetIntFunc(&s, 20) // func require correct type: use pointer, cannot use MyStruct
	p.SetInt(3)        // use pointer
	p.Print()          // can use MyStruct or pointer for method receiver
	Print(s)           // must use MyStruct cannot use pointer: Print(&s) does not compile
	SetIntFunc(p, 30)  // use pointer
}
