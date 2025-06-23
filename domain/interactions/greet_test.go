package interactions_test

import (
	"testing"

	gsg "github.com/fjahn78/go-specs-greet/domain/interactions"
	spec "github.com/fjahn78/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	spec.GreetSpecification(
		t,
		spec.GreetAdapter(gsg.Greet),
	)
}
