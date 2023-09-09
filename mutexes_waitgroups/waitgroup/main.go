package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

// go run .
func main() {
	processWords()
	respectOrder()
}

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func processWords() {
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"zeta",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
}

func respectOrder() {
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()
}
