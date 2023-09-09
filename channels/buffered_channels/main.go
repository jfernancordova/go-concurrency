package main

import (
	"fmt"
	"time"
)

func main() {
	// buffered channel
	// it is helpful when you know how many go routines you've launched or
	// we want to limit the amount of work
	ch := make(chan int, 10)

	go listen(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel")
		ch <- i
		fmt.Println("sent", i, "to channel")
	}

	fmt.Println("Done")
	close(ch)
}

func listen(ch chan int) {
	for {
		i := <-ch
		fmt.Println("got", i, "from channel")

		//simulation
		time.Sleep(1 * time.Second)
	}
}
