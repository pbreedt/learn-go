package types

import "fmt"

func makeNew() {
	mMap := make(map[string]string) // returns initialized map
	pMap := map[string]string{}     // returns initialized map
	nMap := new(map[string]string)  // returns pointer to uninitialized map

	fmt.Printf("Make: %+v\n", mMap)
	fmt.Printf("Plain: %+v\n", pMap)
	fmt.Printf("New: %+v\n", nMap)

	mMap["a"] = "A"             // use directly
	pMap["a"] = "A"             // use directly
	nMap = &map[string]string{} // without this line: panics with 'assignment to entry in nil map'
	(*nMap)["a"] = "A"          // needs initialization first

	fmt.Printf("Make: %+v\n", mMap)
	fmt.Printf("Plain: %+v\n", pMap)
	fmt.Printf("New: %+v\n", nMap)

	miInt := make([]int, 1) // returns initialized slice
	muInt := make([]int, 0) // returns uninitialized slice
	piInt := []int{0}       // returns initialized slice
	puInt := []int{}        // returns uninitialized slice
	nInt := new([]int)      // returns pointer to uninitialized slice

	fmt.Printf("Make: %+v\n", miInt)
	fmt.Printf("Make: %+v\n", muInt)
	fmt.Printf("Plain: %+v\n", piInt)
	fmt.Printf("Plain: %+v\n", puInt)
	fmt.Printf("New: %+v\n", nInt)

	miInt[0] = 1             // use directly
	muInt = append(muInt, 0) // can append, but not assign
	muInt[0] = 1             // needs initialization first
	piInt[0] = 1             // use directly
	puInt = append(puInt, 0) // can append, but not assign
	puInt[0] = 1             // use directly
	*nInt = append(*nInt, 0) // can append, but not assign
	nInt = &[]int{0}         // need to increase capacity
	(*nInt)[0] = 1           // needs initialization first, with either of lines above

	fmt.Printf("Make: %+v\n", miInt)
	fmt.Printf("Make: %+v\n", muInt)
	fmt.Printf("Plain: %+v\n", piInt)
	fmt.Printf("Plain: %+v\n", puInt)
	fmt.Printf("New: %+v\n", nInt)
}
