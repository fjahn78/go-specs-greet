package httpserver

import (
	"fmt"
	"net/http"

	gsg "github.com/fjahn78/go-specs-greet/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, gsg.Greet(name))
}
