package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("test")
	wg.Wait()

	if msg != "test" {
		t.Error("Expected to find test, but it is not there")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "test"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "test") {
		t.Error("Expected to find test, but it is not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, world!") {
		t.Error("Expected to find Hello, world!, but it is not there")
	}
}
