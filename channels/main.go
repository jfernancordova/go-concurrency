package main

import (
	"fmt"
	"strings"
)

// go run .
func main() {
	// two channels: ping is what we send to, pong is what comes back.
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press Enter (enter Q to quit)")
	for {
		// print a prompt
		fmt.Print("-> ")
		// user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for a response
		res := <-pong
		fmt.Println("Response:", res)
	}

	fmt.Println("All done, closing channels")
	close(ping)
	close(pong)
}

// To specify a channel to be a sender or a receiver.
// ping is the receiver and pong is the sender.
// It helps to prevent accidents.
func shout(ping <-chan string, pong chan<- string) {
	for {
		// read from the ping channel. Goroutines waits here -- it blocks until
		// something is received on this channel.
		s, ok := <-ping
		// this is an easy way to make sure that the channel is in fact not closed
		if !ok {
			fmt.Println("channel is closed.")
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}
