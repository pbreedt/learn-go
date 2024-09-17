package dfer

func basics() {
	order()
}

func order() {
	for i := 0; i < 5; i++ {
		// defers executed in LIFO order
		defer println(i)
	}
}
