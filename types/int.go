package types

import "fmt"

func overflow() int8 {
	var max int8 = 127
	// overflow, results in -128
	fmt.Println("int8: max=", max, ", max + 1=", max+1)
	return max + 1
}
