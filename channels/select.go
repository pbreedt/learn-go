package channels

import (
	"fmt"
	"time"
)

func doSelect() {
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		for j := 1; j < 10; j++ {
			fmt.Printf("Sending %d\n", j)
			messages <- fmt.Sprintf("j:%d", j)
			signals <- j > 5
			time.Sleep(1 * time.Second)
		}
		close(messages)
		close(signals)
	}()

	i := 0
	for {
		i++
		fmt.Println("loop top", i)
		select {
		case msg, ok := <-messages:
			if !ok {
				fmt.Println("Messages done")
			}
			fmt.Println("received message", msg)
		case sig, ok := <-signals:
			if !ok {
				fmt.Println("Signals done")
				signals = nil // makes future selects fail for this case, thus only reading from messages
			}
			fmt.Println("received signal", sig)
		}
		fmt.Println("loop end", i)
		if i > 30 {
			return
		}
	}
}
