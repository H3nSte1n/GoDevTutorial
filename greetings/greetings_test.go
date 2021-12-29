package greetings

import (
	"testing"
	"regexp"
)

func TestGreetingName(t *testing.T) {
	name := "World"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Greeting(name)

	if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Greeting("%q") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestGreetingEmpty(t *testing.T) {
	msg, err := Greeting("")
	if msg != "" || err == nil {
			t.Fatalf(`Greeting("") = %q, %v, want "", error`, msg, err)
	}
}

func TestRandomGreetingsName(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie"}
	messages, err := RandomGreetings(names)

	if err != nil {
			t.Fatalf("RandomGreetings(%#v) = %v, want nil", names, err)
	}

	for _, name := range names {
		want := regexp.MustCompile(name)
		if msg, ok := messages[name]; !ok || !want.MatchString(msg) {
				t.Fatalf(`RandomGreetings(%#v) = %#v, want match for %#q`, names, messages, want)
		}
	}
}

func TestRandomGreetingsEmptyArray(t *testing.T) {
	messages, err := RandomGreetings([]string{})

	if err == nil {
			t.Fatalf("RandomGreetings([]string{}) = %v, want error", err)
	}

	if messages != nil {
			t.Fatalf("RandomGreetings([]string{}) = %#v, want nil", messages)
	}
}

func TestRandomGreetingsEmptyArraySlot(t *testing.T) {
	names := []string{"Alice", "", "Charlie"}
	messages, err := RandomGreetings(names)

	if err == nil {
			t.Fatalf("RandomGreetings(%#v) = %v, want error", names, err)
	}

	if messages != nil {
			t.Fatalf("RandomGreetings(%#v) = %#v, want nil", names, messages)
	}
}