package lib

import (
	"foxyfy/internal/lib"
	"testing"
)

func TestGreeting(t *testing.T) {
	greeting := lib.Greet("Mika")
	if greeting != "hi Mika" {
		t.Fatalf("expected: hello Mika, actual %s", greeting)
	}
}
