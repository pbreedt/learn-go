package testing

import "fmt"

func basics() {
	m := make(map[string]int)

	m["one"] = 1
	fmt.Println(m["one"])

	m["one"] = 2
	fmt.Println(m["one"])

	delete(m, "one") // we can delete from map (not from slice)
	fmt.Println(m["one"])

	v, ok := m["one"]
	fmt.Println("The value:", v, "Present?", ok)

	// also works with uninitialized map
	var mm map[string]string
	s, okk := mm["one"]
	fmt.Println("The value:", s, "Present?", okk, "len(map)", len(mm))
	delete(mm, "one") // no panic
	// mm["two"] = "2" // PANICS - nil map

}
