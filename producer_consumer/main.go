package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numberOfPizzas = 10

var successPizzas, failedPizzas, total int

// producer is a type for structs that holds two channels: one for pizzas, with all
// information for a given pizza order including whether it was made successfully
// and another to handle end of processing (when we quit the channel)
type producer struct {
	// Pizza order
	data chan order
	// Notice of not attempting pizzas for whatever reason
	quit chan chan error
}

// order is a type for structs that describes a given pizza order. It has the order number
// a message indicating what happened to the order, and a boolean
// indicating if the order was successfully completed.
type order struct {
	number  int
	message string
	success bool
}

// This program tackles the producer-consumer problem based on a pizzeria.
// go run -race .
func main() {
	// seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// print out a message
	fmt.Println("The pizzeria is open!")
	fmt.Printf("---------------------\n")

	// create a producer
	job := &producer{
		data: make(chan order),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(job)

	// create and run consumer
	for i := range job.data {
		if i.number <= numberOfPizzas {
			if i.success {
				fmt.Printf("Order #%d is out for delivery!\n", i.number)
			} else {
				fmt.Printf("The customer is really mad!\n")
			}
		} else {
			fmt.Println("Done making pizza")
			err := job.close()
			if err != nil {
				fmt.Println("Error closing channel!", err)
			}
		}
	}

	// print out the ending message
	fmt.Printf("---------------------\n")
	fmt.Println("Done for the day.")

	fmt.Printf("We made %d pizzas, but failed to make %d, with %d attempts in total.\n", successPizzas, failedPizzas, total)

	if failedPizzas > 9 {
		fmt.Println("It was an awful day...")
	} else {
		fmt.Println("It was a great day!!!")
	}
}

// pizzeria is a goroutine that runs in the background and
// calls makePizza to try to make one order each time it iterates
// through the loop. It executes until it receives something on the quit
// channel. The quit channel does not receive anything until the consumer
// sends it (when the number of order is greater than or equal the constant NumberOfPizza).
func pizzeria(maker *producer) {
	// keep tack of which pizza we are making
	var i = 0

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.number
			select {
			// trying to make a pizza (we send something to the channel)
			case maker.data <- *currentPizza:
			case quitChan := <-maker.quit:
				// close channels
				close(maker.data)
				close(quitChan)
				return
			}
		}
	}
}

// close is simply a method of closing the channel when we are done with it.
func (p *producer) close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

// makePizza attempts to make a pizza, we generate a random number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(orderNumber int) *order {
	orderNumber++
	if orderNumber <= numberOfPizzas {
		delay := rand.Intn(5)
		fmt.Printf("Received order #%d!\n", orderNumber)

		// timer to make the pizza between 1 and 12
		timer := rand.Intn(12) + 1
		msg := ""
		success := false

		if timer < 5 {
			failedPizzas++
		} else {
			successPizzas++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds ...\n", orderNumber, delay)

		// delay for a bit in seconds
		time.Sleep(time.Duration(delay) * time.Second)

		if timer <= 2 {
			msg = fmt.Sprintf(" *** We ran out of ingredients for pizza #%d! ***", orderNumber)
		} else if timer <= 4 {
			msg = fmt.Sprintf(" *** The cook quit while making pizza #%d! ***", orderNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", orderNumber)
		}

		return &order{
			number:  orderNumber,
			message: msg,
			success: success,
		}
	}

	return &order{
		number: orderNumber,
	}
}
