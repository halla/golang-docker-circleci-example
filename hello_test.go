package main

import "testing"

// test function names begin with uppercase Test
func TestMessage(t *testing.T) {
	project := new(Project)
	project.PushedAt = "now"
	project.Name = "TestProject"
	msg := Message(project)
	if msg != "<p>TestProject: Latest commit: now</p>" {
		t.Errorf("Incorrect message: %s", msg)
	}
}
