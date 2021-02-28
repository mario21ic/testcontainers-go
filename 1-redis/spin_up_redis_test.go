package test

import (
  "testing"
  "context"

  "github.com/testcontainers/testcontainers-go"
  "github.com/testcontainers/testcontainers-go/wait"
)

func TestWithRedis(t *testing.T) {
    ctx := context.Background()
    req := testcontainers.ContainerRequest{
        Image:        "redis:latest",
        ExposedPorts: []string{"6379/tcp"},
        WaitingFor:   wait.ForLog("Ready to accept connections"),
    }
    redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: req,
        Started:          true,
    })
    if err != nil {
        t.Error(err)
    }
    defer redisC.Terminate(ctx)
}
