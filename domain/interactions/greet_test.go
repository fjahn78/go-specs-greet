package interactions_test

import (
	"testing"

	"github.com/fjahn78/go-specs-greet/domain/interactions"
	spec "github.com/fjahn78/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	t.Run("a name is given", func(t *testing.T) {
		spec.GreetSpecification(
			t,
			spec.GreetAdapter(interactions.Greet),
		)
	})
	// t.Run()
}
