package switches

import "fmt"

func WithExpression() {
	value := "something"

	switch value {
	case "something":
		fmt.Println("it's something")
	default:
		fmt.Println("it's probably nothing")
	}
}

func WithExtendedExpression() {
	switch value := "something"; value {
	case "something":
		fmt.Println("it's something")
	default:
		fmt.Println("it's probably nothing")
	}
}

func WithoutExpression() {
	i := 10
	switch {
	case i > 10:
		fmt.Println("Ohh, big!")
	case i < 10:
		fmt.Println("Mmm, small.")
	default:
		fmt.Println("Just right.")
	}
}
