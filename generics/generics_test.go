package generics

import (
	"reflect"
	"testing"
)

func TestSumMapIntsOrFloatsExplicit(t *testing.T) {
	i := SumMapIntsOrFloats[string, int64](map[string]int64{"first": 34, "second": 12})
	want := reflect.Int64.String()
	got := reflect.TypeOf(i).Kind().String()
	if want != got {
		t.Errorf("incorrect type. wanted:%s, got:%s", want, got)
	}
	if i != 46 {
		t.Errorf("incorrect result. wanted:46, got:%d", i)
	}

	f := SumMapIntsOrFloats[string, float64](map[string]float64{"first": 35.98, "second": 26.99})
	want = reflect.Float64.String()
	got = reflect.TypeOf(f).Kind().String()
	if want != got {
		t.Errorf("incorrect type. wanted:%s, got:%s", want, got)
	}
	if f != 62.97 {
		t.Errorf("incorrect result. wanted:62.97, got:%f", f)
	}
}

func TestSumMapIntsOrFloatsInfered(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	i := SumMapIntsOrFloats(ints)
	want := reflect.Int64.String()
	got := reflect.TypeOf(i).Kind().String()
	if want != got {
		t.Errorf("incorrect type. wanted:%s, got:%s", want, got)
	}
	if i != 46 {
		t.Errorf("incorrect result. wanted:46, got:%d", i)
	}

	f := SumMapIntsOrFloats(floats)
	want = reflect.Float64.String()
	got = reflect.TypeOf(f).Kind().String()
	if want != got {
		t.Errorf("incorrect type. wanted:%s, got:%s", want, got)
	}
	if f != 62.97 {
		t.Errorf("incorrect result. wanted:62.97, got:%f", f)
	}
}

func TestSumNumbers(t *testing.T) {
	// int not int64
	i := SumNumbers([]int{34, 12})
	want := reflect.Int.String()
	got := reflect.TypeOf(i).Kind().String()
	if want != got {
		t.Errorf("incorrect type. wanted:%s, got:%s", want, got)
	}
	if i != 46 {
		t.Errorf("incorrect result. wanted:46, got:%d", i)
	}

	f := SumNumbers([]float64{35.98, 26.99})
	want = reflect.Float64.String()
	got = reflect.TypeOf(f).Kind().String()
	if want != got {
		t.Errorf("incorrect type. wanted:%s, got:%s", want, got)
	}
	if f != 62.97 {
		t.Errorf("incorrect result. wanted:62.97, got:%f", f)
	}
}
