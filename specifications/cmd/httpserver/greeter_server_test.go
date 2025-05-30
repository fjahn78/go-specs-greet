package main_test

import (
	"testing"

	go_specs_greet "lsscmp.fin-rlp.local/ZD0436admin/go-specs-greet"
	"lsscmp.fin-rlp.local/ZD0436admin/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
