package types

import (
	"reflect"
	"testing"
)

func TestBasics(t *testing.T) {
	basics()
}

func TestReflect(t *testing.T) {
	var mystr string
	type strWrapper string

	testcases := []struct {
		given   interface{}
		expType string
		expKind reflect.Kind
	}{
		{"hello", "string", reflect.String},
		{mystr, "string", reflect.String},
		{strWrapper("wrapped string"), "types.strWrapper", reflect.String},
	}

	for _, tc := range testcases {
		typ := GetType(tc.given)
		knd := GetKind(tc.given)
		if typ != tc.expType {
			t.Errorf("incorrect type. wanted:%s, got:%s", tc.expType, typ)
		}
		if knd != tc.expKind {
			t.Errorf("incorrect kind. wanted:%s, got:%s", tc.expKind, knd)
		}
	}

}
