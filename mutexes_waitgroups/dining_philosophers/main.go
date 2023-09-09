package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously, since there are five philosophers and five forks.
//
// This is a simple implementation of Dijkstra's solution to the "Dining
// Philosophers" dilemma.

type philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Kant", leftFork: 3, rightFork: 4},
}

var hunger = 3                  // how many times a philosopher eats
var eatTime = 1 * time.Second   // how long it takes to eatTime
var thinkTime = 3 * time.Second // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out

var orderFinished []string // the order in which philosophers finish dining and leave
var orderMutex sync.Mutex  // to control the name manipulation when appends the name

// go run -race .
func main() {
	// welcome message
	fmt.Println("Dining philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	// finished message
	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)
	fmt.Printf("The order in which every philosopher ended his meal was: %s.\n", strings.Join(orderFinished, ", "))
}

func dine() {
	l := len(philosophers)
	wg := &sync.WaitGroup{}
	wg.Add(l)

	// if a philosopher is seated
	seated := &sync.WaitGroup{}
	seated.Add(l)

	// forks is a map of all 5 forks.
	// it needs the pointer because once you've created a mutex you're never supposed to copy it.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < l; i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal by iterating through our slice of philosophers
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go dinning(philosophers[i], wg, forks, seated)
	}

	// Pause the program execution until all five goroutine are actually done.
	wg.Wait()
}

// the function fired off as a goroutine for each of our philosophers. It takes one
// philosopher, our WaitGroup to determine when everyone is done, a map containing
// the mutexes for every fork on the table, and a WaitGroup used to pause execution of every instance
// of this goroutine until everyone is seated at the table.
func dinning(p philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n\n", p.name)
	seated.Done()

	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		// this goroutine is paused

		// logical race condition: all the things lock! It would never detect a race condition
		// we need to make sure that we don't have two philosopher taking the wrong fork that they will wait endlessly
		if p.leftFork > p.rightFork {
			forks[p.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork. \n", p.name)

			forks[p.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork. \n", p.name)
		} else {
			forks[p.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork. \n", p.name)

			forks[p.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork. \n", p.name)
		}

		fmt.Printf("\t%s is eating. \n", p.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking. \n", p.name)
		time.Sleep(thinkTime)

		forks[p.leftFork].Unlock()
		forks[p.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", p.name)
	}

	fmt.Println(p.name, "is satisfied.")
	fmt.Println(p.name, "left the table.")

	orderMutex.Lock()
	orderFinished = append(orderFinished, p.name)
	orderMutex.Unlock()
}
