package main_test

import (
	"fmt"
	"testing"

	"github.com/fjahn78/go-specs-greet/adapters"
	"github.com/fjahn78/go-specs-greet/adapters/grpcserver"
	"github.com/fjahn78/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpc_server/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)
	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)
}
