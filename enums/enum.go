package enums

// See promising implementation using generics: https://github.com/orsinium-labs/enum

type Weekday int

// Flawed enum: Weekday(int) is valid for any int
// Also, default value of Weekday is 0, which is Monday but might be unintentional
const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Season int

// Flawed enum: Season(int) is valid for any int
const (
	Undefined Season = iota
	Summer           // 1
	Winter           // 2
	Spring           // 3
	Fall             // 4
)

// IsValid ensures that the enum value is valid, although it is not enforced by the compiler
func (s Season) IsValid() bool {
	return int(s) >= 0 && int(s) <= 4
}
func (s Season) String() string {
	return [...]string{"Undefined", "Summer", "Winter", "Spring", "Fall"}[s]
}

type UserRole int

// Bitwise enum, allows for bit operations:
// var role UserRole = Member | Moderator // 6
const (
	Guest     UserRole = 1 << iota // 1
	Member                         // 2
	Moderator                      // 4
	Admin                          // 8
)

func (ur UserRole) IsValid() bool {
	var maxRole UserRole = (Admin | Member | Moderator | Guest)
	return int(ur) >= 1 && int(ur) <= int(maxRole)
}

type Pet string

const (
	Dog  Pet = "dog"
	Cat  Pet = "cat"
	Fish Pet = "fish"
	Bird Pet = "bird"
)

// Nothing forces user to use this constructor
func NewPet(s string) *Pet {
	if s == "dog" || s == "cat" || s == "fish" || s == "bird" {
		p := Pet(s)
		return &p
	}
	return nil
}
