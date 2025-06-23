package go_specs_greet_test

import (
	"testing"

	gsg "github.com/fjahn78/go-specs-greet"
	spec "github.com/fjahn78/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	spec.GreetSpecification(
		t, 
		spec.GreetAdapter(gsg.Greet),
	)
}
