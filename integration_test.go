package main

import (
	"context"
	"log"
	"testing"
	"time"

	"go.uber.org/fx"
)

func TestLocalGorm(t *testing.T) {
	app := fx.New(
		createIntegrationTestConfigurationGorm(),
		createIntegrationTestInvocation(t),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
}

func TestClientServer(t *testing.T) {
	serverApp := fx.New(
		createIntegrationTestConfigurationServer(),
		createIntegrationTestServerInvocation(),
	)

	started := make(chan bool)
	go func() {
		startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := serverApp.Start(startCtx); err != nil {
			log.Fatal(err)
		}
		started <- true
	}()

	select {
	case <-started:
	}

	fx.New(
		createIntegrationTestConfigurationClient(),
		createIntegrationTestInvocation(t),
	)
}

func TestRoundRobin(t *testing.T) {
	app := fx.New(
		createIntegrationTestConfigurationRoundRobin(),
		createIntegrationTestRoundRobinInvocation(t),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
}
