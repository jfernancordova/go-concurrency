package main

import (
	"sync"
	"testing"
)

// go test -race .
func Test_updateMessage(t *testing.T) {
	msg = "Test"
	var mutex sync.Mutex

	wg.Add(1)
	go updateMessage("Second", &mutex)
	wg.Wait()

	if msg != "Second" {
		t.Error("Incorrect value in msg")
	}
}
