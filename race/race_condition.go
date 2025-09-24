package race

import "fmt"

func Race() {
	i := 0
	go func() {
		i++ // write
	}()
	i++
	fmt.Println(i) // concurrent read
}

func NoRace() {
	c := make(chan int)
	go func() {
		locali := 0
		locali++ // write
		c <- locali
	}()

	i := <-c
	i++
	fmt.Println(i) // no concurrent read
}
