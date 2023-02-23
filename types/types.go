package types

import "fmt"

func basics() {
	typeSwitch(21)
	typeSwitch("hello")
	typeSwitch(true)

	typeSwitchWithValue(21)
	typeSwitchWithValue("hello")
	typeSwitchWithValue(true)

	typeAssert()
	typeCast()
}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("i is int but 'cast': %T (can't do ops on it i*2 does not compile)\n", i)
	case string:
		fmt.Printf("i is string but 'cast': %T (can't do string ops on it len(i) does not compile)\n", i)
	default:
		fmt.Printf("i is unknown type: %T\n", i)
	}
}

func typeSwitchWithValue(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is int. (x2=%d)\n", v, v*2)
	case string:
		fmt.Printf("%q is string. (len=%d)\n", v, len(v))
	default:
		fmt.Printf("unknown type %T!\n", v)
	}
}

func typeAssert() {
	var i interface{} = "hello" // asserts require interface

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok) // zero value of f, ok = false

	// f = i.(float64) // panic
	// fmt.Println(f)
}

func typeCast() {
	var i int32 = 4 // cast require compatible concrete types ()
	fmt.Printf("(%v, %T)\n", i, i)

	i64 := int64(i)
	fmt.Printf("(%v, %T)\n", i64, i64)

	var intf I
	s := S{"aaa"}
	s.M()
	s.M2()
	intf = I(s) // can cast a struct to an interface (if struct implements interface)
	intf.M()
	// intf.M2() // interface does not have M2 method

	intWrap := IntWrapper(3) // cast int to IntWrapper
	intVal := int(intWrap)   // cast IntWrapper to int
	fmt.Printf("int:%d | IntWrapper:%d\n", intVal, intWrap)
}

type IntWrapper int

type I interface {
	M()
}

type S struct {
	val string
}

func (s S) M() {
	fmt.Printf("M: (%v, %T)\n", s, s)
}

func (s S) M2() {
	fmt.Printf("M2: (%v, %T)\n", s, s)
}
