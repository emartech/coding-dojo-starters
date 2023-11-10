package greeter_test

import (
	"testing"

	"github.com/emartech/coding-dojo-starters/pkg/greeter"
)

func TestGivenNoparamsReturnsGreeting(t *testing.T) {
	want := "Hello Dojoers!" // go "standard" for expected

	got := greeter.Hello() // go "standard" for actual

	if got != want {
		t.Errorf(`got "%s", want "%s"`, got, want)
	}
}
