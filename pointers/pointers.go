package pointers

import "fmt"

type MyStruct struct {
	strVal string
	intVal int
}

func basics() {
	i := 10
	p := &i         // pointer to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p)  // pointer value
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	s := MyStruct{"aa", 10}
	pp := &s                  // pointer to struct
	fmt.Println((*pp).strVal) // read property through the pointer wihg dereference
	fmt.Println(pp.strVal)    // read property through the pointer (no deref * required)

	m := new(MyStruct) // new return pointer
	fmt.Printf("%T %v\n", m, m)
}
