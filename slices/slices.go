package slices

import "fmt"

/* TODO:
- Add Generics
- Add IndexOf
- Add RemoveAll*
*/

func basics() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // Last index excluded
	b := a          //copy by assignment, ref same underlying array
	c := make([]string, len(b))
	copy(c, b) // copy after 'make', new underlying array
	fmt.Println("1>", a, b, c)

	b[0] = "XXX" // change to slice (even copy of slice) affects underlying array as well
	c[1] = "OOO" // change to new slice (with make) does not affect underlying array as well
	fmt.Println("2>", a, b, c)
	fmt.Println("3>", names)

	aa := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("4>", aa[0:10]) // all the same
	fmt.Println("5>", aa[:10])  // all the same
	fmt.Println("6>", aa[0:])   // all the same
	fmt.Println("7>", aa[:])    // all the same
}

func lenAndCap() {
	s := []int{0, 1, 2, 3, 4, 5}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop from the front reduces capacity
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Copy() {
	data := []int{1, 2, 3, 4, 5}
	cpy := make([]int, len(data))
	// cpy = append(cpy, data...) // ERROR: make already initializes 5 elements with default values, so [0,0,0,0,0], append just adds to this array
	copy(cpy, data) // CORRECT: copies data to new slice. NOTE: min length of len(data) and len(cpy) will be used
}

func RemoveIndexString(s []string, i int) []string {
	newSlice := make([]string, 0)           // create new empty slice
	newSlice = append(newSlice, s[:i]...)   // add all elements - up to element i
	newSlice = append(newSlice, s[i+1:]...) // add rest of elements - from 1 after i
	return newSlice
}

func RemoveFirstString(s []string, se string) []string {
	for i, elem := range s {
		if elem == se {
			return RemoveIndexString(s, i)
		}
	}

	return s
}

func ContainsString(txs []string, tx string) bool {
	for _, t := range txs {
		if t == tx {
			return true
		}
	}
	return false
}

// Generic impl
// func Contains[T comparable](s []T, e T) bool {
// 	for _, v = range s {
// 		if v == e {
// 			return true
// 		}
// 	}
// 	return false
// }
