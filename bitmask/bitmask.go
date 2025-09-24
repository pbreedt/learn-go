package bitmask

import (
	"fmt"
)

type Bits uint8

const (
	RightBit Bits = 1 << iota
	MiddleBit
	LeftBit
)

func Set(b, flag Bits) Bits    { return b | flag }       // OR (for each bit that is set in flag, the corresponding bit in b will be set to 1)
func Clear(b, flag Bits) Bits  { return b &^ flag }      // AND NOT (for each bit that is set in flag, the corresponding bit in b will be set to 0)
func Toggle(b, flag Bits) Bits { return b ^ flag }       // XOR (for each bit that is set in flag, the corresponding bit in b will be flipped)
func Has(b, flag Bits) bool    { return b&flag != 0 }    // AND (checks if any of the bits in flag are set in b)
func Any(b Bits) bool          { return b != 0 }         // checks if any bit is set in b
func All(b, flag Bits) bool    { return b&flag == flag } // checks if all bits in flag are set in b
func None(b, flag Bits) bool   { return b&flag == 0 }    // checks if none of the bits in flag are set in b
func IsOnly(b, flag Bits) bool { return b == flag }      // checks if only the bits in flag are set in b

func BitsToString(b Bits) string {
	flags := []string{"RightBit", "MiddleBit", "LeftBit"}
	result := ""

	// Iterate over flags backwards so that order matches position Left|Middle|Right
	for i := len(flags); i >= 0; i-- {
		//      001<<i: 001(i=0), 010(i=1), 100(i=2)
		if Has(b, 1<<i) {
			if result != "" {
				result += "|"
			}
			result += flags[i]
		}
	}
	result += fmt.Sprintf("(%03b)(%d)", b, b)
	return result
}
