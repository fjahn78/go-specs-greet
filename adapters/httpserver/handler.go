package httpserver

import (
	"fmt"
	"log"
	"net/http"

	gsg "github.com/fjahn78/go-specs-greet/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if _, err := fmt.Fprint(w, gsg.Greet(name)); err!= nil {
		log.Fatal(err)
	}
}
