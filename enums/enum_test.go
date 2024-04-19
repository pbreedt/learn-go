package enums

import (
	"reflect"
	"testing"
)

type TestCase[T1 any, T2 any] struct {
	Given   T1
	Want    T2
	IsValid bool
}

func TestWeekdayEnum(t *testing.T) {
	test := []TestCase[int, Weekday]{
		{Given: 0, Want: Monday},
		{Given: 1, Want: Tuesday},
		{Given: 2, Want: Wednesday},
		{Given: 3, Want: Thursday},
		{Given: 4, Want: Friday},
		{Given: 5, Want: Saturday},
		{Given: 6, Want: Sunday},
	}
	for _, tc := range test {
		got := Weekday(tc.Given)
		if tc.Want != got {
			t.Errorf("incorrect result. wanted:%d, got:%d", tc.Want, got)
		}
	}
}

func TestSeasonEnum(t *testing.T) {
	test := []TestCase[int, Season]{
		{Given: 0, Want: Undefined, IsValid: true},
		{Given: 1, Want: Summer, IsValid: true},
		{Given: 2, Want: Winter, IsValid: true},
		{Given: 3, Want: Spring, IsValid: true},
		{Given: 4, Want: Fall, IsValid: true},
		{Given: 5, Want: Season(5), IsValid: false},
	}
	for _, tc := range test {
		got := Season(tc.Given)
		if tc.Want != got {
			t.Errorf("incorrect result. wanted:%d, got:%d", tc.Want, got)
		}
		if tc.IsValid != got.IsValid() {
			t.Errorf("invalid result. wanted:%t, got:%t", tc.IsValid, got.IsValid())
		}
		if got.IsValid() {
			t.Logf("valid season: %s", got)
		}
	}
}

func TestUserRoleEnum(t *testing.T) {
	test := []TestCase[int, int]{
		{Given: 0, Want: 0, IsValid: false},
		{Given: 1, Want: int(Guest), IsValid: true},
		{Given: 2, Want: int(Member), IsValid: true},
		{Given: 3, Want: int(Guest | Member), IsValid: true},
		{Given: 4, Want: int(Moderator), IsValid: true},
		{Given: 5, Want: int(Guest | Moderator), IsValid: true},
		{Given: 6, Want: int(Member | Moderator), IsValid: true},
		{Given: 7, Want: int(Guest | Member | Moderator), IsValid: true},
		{Given: 8, Want: int(Admin), IsValid: true},
		{Given: 9, Want: int(Guest | Admin), IsValid: true},
		{Given: 10, Want: int(Member | Admin), IsValid: true},
		{Given: 11, Want: int(Guest | Member | Admin), IsValid: true},
		{Given: 12, Want: int(Moderator | Admin), IsValid: true},
		{Given: 13, Want: int(Guest | Moderator | Admin), IsValid: true},
		{Given: 14, Want: int(Member | Moderator | Admin), IsValid: true},
		{Given: 15, Want: int(Guest | Member | Moderator | Admin), IsValid: true},
	}
	for _, tc := range test {
		got := UserRole(tc.Given)
		if tc.Want != int(got) {
			t.Errorf("incorrect result. wanted:%d, got:%d", tc.Want, got)
		}
		if tc.IsValid != got.IsValid() {
			t.Errorf("invalid result for value %d. wanted:%t, got:%t", tc.Given, tc.IsValid, got.IsValid())
		}
	}
}

func TestPetEnum(t *testing.T) {
	test := []TestCase[string, any]{
		{Given: "cat", Want: Cat, IsValid: true},
		{Given: "dog", Want: Dog, IsValid: true},
		{Given: "fish", Want: Fish, IsValid: true},
		{Given: "bird", Want: Bird, IsValid: true},
		{Given: "horse", Want: nil, IsValid: false},
	}
	for _, tc := range test {
		got := NewPet(tc.Given)
		if !tc.IsValid {
			// "horse" is not a valid pet, so NewPet should return nil
			if got != nil {
				t.Errorf("incorrect result. wanted:%v, got:%v", tc.Want, got)
			}
		} else {
			if reflect.TypeOf(*got).Name() != reflect.TypeOf(tc.Want).Name() {
				t.Errorf("incorrect result. wanted:%v, got:%v", tc.Want, got)
			}
		}
	}
}
