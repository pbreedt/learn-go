package types

import (
	"fmt"
	"reflect"
)

func basics() {
	typeSwitch(21)
	typeSwitch("hello")
	typeSwitch(true)

	typeSwitchWithValue(21)
	typeSwitchWithValue("hello")
	typeSwitchWithValue(true)

	typeAssert()
	typeCast()

	var mystr string
	fmt.Println("## Type of empty string:")
	typeReflect(mystr)
	fmt.Println("\n## Type of string with value:")
	typeReflect("my string value")
	type strWrapper string // Type = types.strWrapper, Kind = string
	fmt.Println("\n## Type of string wrapped as strWrapper type:")
	typeReflect(strWrapper("wrapped string"))
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

func typeReflect(var1 interface{}) {
	fmt.Printf("Printf %%T of var1: %T\n", var1)
	fmt.Println("TypeOf var1:", reflect.TypeOf(var1))
	fmt.Println("TypeOf.Kind var1:", reflect.TypeOf(var1).Kind())
	fmt.Println("ValueOf var1:", reflect.ValueOf(var1))
	fmt.Println("ValueOf.Kind var1:", reflect.ValueOf(var1).Kind())

	// value is the actual value/contents of the variable - not relevant to type
	if reflect.ValueOf(var1) == reflect.ValueOf("") {
		fmt.Println("ValueOf is empty string")
	} else {
		fmt.Println("ValueOf is NOT empty string, it is", reflect.ValueOf(var1))
	}

	if reflect.TypeOf(var1).String() == "string" {
		fmt.Println("TypeOf.String() is 'string'")
	} else {
		fmt.Println("TypeOf.String() is NOT 'string', it is", reflect.TypeOf(var1).String())
	}

	if reflect.TypeOf(var1).Kind() == reflect.String {
		fmt.Println("TypeOf.Kind() is reflect.String")
	} else {
		fmt.Println("TypeOf.Kind() is NOT reflect.String, it is", reflect.TypeOf(var1).Kind())
	}

	if reflect.TypeOf(var1).Kind().String() == "string" {
		fmt.Println("TypeOf.Kind().String() is 'string'")
	} else {
		fmt.Println("TypeOf.Kind().String() is NOT 'string'")
	}
}

func GetType(var1 interface{}) string {
	return reflect.TypeOf(var1).String()
}
func GetKind(var1 interface{}) reflect.Kind {
	return reflect.TypeOf(var1).Kind()
}
