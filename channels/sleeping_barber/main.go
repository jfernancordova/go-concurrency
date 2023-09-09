package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

// go run -race .
func main() {
	// Let's break this rather complex problem

	// seed our random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// print welcome message
	fmt.Println("The Sleeping Barber Problem")
	fmt.Println("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create ds to represent the barbershop
	shop := barberShop{
		shopCapacity:    seatingCapacity,
		hairCutDuration: cutDuration,
		numberOfBarbers: 0,
		clientsChan:     clientChan,
		barbersDoneChan: doneChan,
		open:            true,
	}

	fmt.Println("The shop is open for the day!")

	// add barbers
	shop.addBarber("Fernando")
	shop.addBarber("José")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		// block until time open passes
		// received value ignored
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1
	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			// this is the point when a client arrives at random intervals
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
