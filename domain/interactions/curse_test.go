package interactions_test

import (
	"testing"

	"github.com/fjahn78/go-specs-greet/domain/interactions"
	spec "github.com/fjahn78/go-specs-greet/specifications"
)

func TestCurse(t *testing.T) {
	spec.CurseSpecification(
		t,
		spec.CurseAdapter(interactions.Curse),
	)
}
