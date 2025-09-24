package bitmask

import (
	"testing"
)

func TestBitmaskSet(t *testing.T) {
	var b Bits

	// Initially, no flags are set
	if Any(b) {
		t.Error("Expected no flags to be set")
	}

	// Set RightBit
	b = Set(b, RightBit)
	if !Has(b, RightBit) {
		t.Error("Expected RightBit to be set")
	}
	if !Any(b) {
		t.Error("Expected some flags to be set")
	}
	if IsOnly(b, RightBit) == false {
		t.Error("Expected only RightBit to be set")
	}
}

func TestBitmaskClear(t *testing.T) {
	var b Bits

	// Set RightBit and MiddleBit
	b = Set(b, RightBit)
	b = Set(b, MiddleBit)

	// Clear RightBit
	b = Clear(b, RightBit)
	if Has(b, RightBit) {
		t.Error("Expected RightBit to be cleared")
	}
	if !Has(b, MiddleBit) {
		t.Error("Expected MiddleBit to still be set")
	}
	if Any(b) == false {
		t.Error("Expected some flags to be set")
	}
	if IsOnly(b, MiddleBit) == false {
		t.Error("Expected only MiddleBit to be set")
	}
}

func TestBitmaskToggle(t *testing.T) {
	var b Bits

	// Toggle RightBit (set it)
	b = Toggle(b, RightBit)
	if !Has(b, RightBit) {
		t.Error("Expected RightBit to be set after toggle")
	}

	// Toggle RightBit again (clear it)
	b = Toggle(b, RightBit)
	if Has(b, RightBit) {
		t.Error("Expected RightBit to be cleared after second toggle")
	}
}

func TestBitmaskAllNone(t *testing.T) {
	var b Bits

	// Set RightBit and MiddleBit
	b = Set(b, RightBit)
	b = Set(b, MiddleBit)

	if !All(b, RightBit|MiddleBit) {
		t.Error("Expected both RightBit and MiddleBit to be set")
	}
	if None(b, LeftBit) == false {
		t.Error("Expected LeftBit to be not set")
	}
}

func TestBitmaskIsOnly(t *testing.T) {
	var b Bits

	// Set RightBit
	b = Set(b, RightBit)
	if !IsOnly(b, RightBit) {
		t.Error("Expected only RightBit to be set")
	}

	// Set MiddleBit as well
	b = Set(b, MiddleBit)
	if IsOnly(b, RightBit) {
		t.Error("Expected more than just RightBit to be set")
	}
}

func TestBitsToString(t *testing.T) {
	var b Bits
	var got string

	// No flags set
	if got = BitsToString(b); got != "(000)(0)" {
		t.Error("Expected '(000)(0)' string for no flags, got:", got)
	}

	// Set RightBit
	b = Set(b, RightBit)
	if got = BitsToString(b); got != "RightBit(001)(1)" {
		t.Error("Expected 'RightBit(001)(1)', got:", got)
	}

	// Set MiddleBit
	b = Set(b, MiddleBit)
	if got = BitsToString(b); got != "MiddleBit|RightBit(011)(3)" {
		t.Error("Expected 'MiddleBit|RightBit(011)(3)', got:", got)
	}

	// Set LeftBit
	b = Set(b, LeftBit)
	if got = BitsToString(b); got != "LeftBit|MiddleBit|RightBit(111)(7)" {
		t.Error("Expected 'LeftBit|MiddleBit|RightBit(111)(7)', got:", got)
	}
}

// func TestSetInLoop(t *testing.T) {
// 	var b Bits
// 	for i := 0; i < 3; i++ {
// 	      001 << i: 001(i=0), 010(i=1), 100(i=2)
// 		b = 1 << i // Set bit i
// 		t.Logf("Set bit %d: %s", i, BitsToString(b))
// 	}
// }
