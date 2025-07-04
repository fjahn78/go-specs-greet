package main_test

import (
	"fmt"
	"testing"

	"github.com/fjahn78/go-specs-greet/adapters"
	"github.com/fjahn78/go-specs-greet/adapters/grpcserver"
	"github.com/fjahn78/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	
	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)
	adapters.StartDockerServer(t, port, "grpc_server")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
