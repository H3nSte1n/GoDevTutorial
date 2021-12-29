package main

import (
	"testing"
	"example.com/greetings"
)

func TestMainCallRandomGreetingsSuccessfully(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie"}
	messages, err := greetings.RandomGreetings(names)

	if err != nil {
		t.Fatalf("RandomGreetings(%#v) = %v, want error", names, err)
	}

	if messages == nil {
		t.Fatalf("RandomGreetings(%#v) = %#v, want nil", names, messages)
	}
}

func TestMainCallRandomGreetingsUnsuccessfully(t *testing.T) {
	names := []string{"Alice", "", "Charlie"}
	messages, err := greetings.RandomGreetings(names)

	if err == nil {
		t.Fatalf("RandomGreetings(%#v) = %v, want error", names, err)
	}

	if messages != nil {
		t.Fatalf("RandomGreetings(%#v) = %#v, want nil", names, messages)
	}
}