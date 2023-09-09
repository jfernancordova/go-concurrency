package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	// exclusive value to access to this variable
	m.Lock()
	msg = s
	m.Unlock()
}

// go run -race
func main() {
	msg = "Test"

	// accessing the data safely
	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Test one", &mutex)
	go updateMessage("Test two", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
