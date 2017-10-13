package main

import "testing"

// test function names begin with uppercase Test
func TestMessage(t *testing.T) {
	project := new(Project)
	project.PushedAt = "now"
	msg := Message(project)
	if msg != "Hello world! Latest commit: now" {
		t.Errorf("Incorrect message: %s", msg)
	}
}
