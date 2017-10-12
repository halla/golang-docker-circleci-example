package main

import "testing"

// test function names begin with uppercase Test
func TestMessage(t *testing.T) {
	msg := Message()
	if msg != "Hello world!" {
		t.Errorf("Incorrect message: %s", msg)
	}
}
