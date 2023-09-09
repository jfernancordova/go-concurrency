package main

import (
	"testing"
	"time"
)

// go test - race .
func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		l := len(orderFinished)
		if l != 5 {
			t.Errorf("expected 5 bug to %d", l)
		}
	}
}

func Test_dineWithDelays(t *testing.T) {
	var delays = []struct {
		name  string
		delay time.Duration
	}{
		{"zero", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, d := range delays {
		orderFinished = []string{}
		eatTime = d.delay
		sleepTime = d.delay
		thinkTime = d.delay

		dine()
		l := len(orderFinished)
		if l != 5 {
			t.Errorf("expected 5 bug to %d", l)
		}
	}
}
