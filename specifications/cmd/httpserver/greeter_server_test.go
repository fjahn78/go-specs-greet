package main_test

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	go_specs_greet "lsscmp.fin-rlp.local/ZD0436admin/go-specs-greet"
	"lsscmp.fin-rlp.local/ZD0436admin/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	req := tc.ContainerRequest{
		FromDockerfile: tc.FromDockerfile{
			Context: "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:		  true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})
	
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
