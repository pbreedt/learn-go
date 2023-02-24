package generics

// SumMapIntsOrFloats only using map to demonstrate comparable
// K comparator required since map requires keys to be 'comparable' / < > == !=
func SumMapIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

func SumNumbers[V Number](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
