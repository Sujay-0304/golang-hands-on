package testable_program

import (
	"testing"
)

func TestSendGreeting(t *testing.T) {
	got := SendGreeting("Sujay")
	want := "Hello, Sujay! Welcome!"

	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}

	t.Log("Integration test passed ")
}