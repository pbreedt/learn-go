package slices

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	basics()
	lenAndCap()
}

func TestRemoveIndexString(t *testing.T) {
	tstSlice := []string{"a", "b", "bb", "c", "d"}
	newSlice := RemoveIndexString(tstSlice, 2)

	if ContainsString(newSlice, "bb") {
		t.Fail()
	}
	fmt.Println("original:", tstSlice, "\n", "new", newSlice)
}

func TestRemoveFirstString(t *testing.T) {
	tstSlice := []string{"a", "b", "bb", "c", "d"}
	newSlice := RemoveFirstString(tstSlice, "bb")

	if ContainsString(newSlice, "bb") {
		t.Fail()
	}
	fmt.Println("original:", tstSlice, "\n", "new", newSlice)
}
